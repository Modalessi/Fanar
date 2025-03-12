package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/Modalessi/iau_resources/database"
	"github.com/Modalessi/iau_resources/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func (s *Storage) StoreResource(ctx context.Context, resource *models.Resource, file io.Reader, contentType string) error {

	err := s.uploadFile(ctx, resource.CourseID.String(), resource.Title, resource.FileExt, file, contentType)
	if err != nil {
		return fmt.Errorf("db error: error uploading file to s3: %s", err)
	}

	resourceKey := s.fileS3Key(resource.CourseID.String(), resource.Title, resource.FileExt)
	resourceDB, err := s.queries.CreateResource(ctx, database.CreateResourceParams{
		CourseID:    resource.CourseID,
		Title:       resource.Title,
		Description: resource.Description,
		FileExt:     resource.FileExt,
		S3Url:       s.fileS3URL(resourceKey),
		Tags:        resource.Tags,
		CreatedBy:   resource.Created_by,
	})

	if err != nil {
		errMsg := fmt.Sprintf("db err: error creating resource in database: %v", err)
		err := s.deleteFile(ctx, resourceKey)
		if err != nil {
			errMsg += fmt.Sprintf("\bdb err: error rolling down s3 upload: %v", err)
		}

		return fmt.Errorf(errMsg)
	}

	resourceURL, err := url.Parse(s.fileS3URL(resourceKey))
	if err != nil {
		return fmt.Errorf("error parsing resource URL: %v", err)
	}

	resource.ID = &resourceDB.ID
	resource.Url = resourceURL

	return nil
}

func (s *Storage) deleteFile(ctx context.Context, key string) error {
	_, err := s.s3.Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.s3.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("aws error: error deleting file from aws %v", err)
	}
	return nil
}

func (s *Storage) uploadFile(ctx context.Context, courseId, fileName, ext string, file io.Reader, contentType string) error {

	key := s.fileS3Key(courseId, fileName, ext)

	_, err := s.s3.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.s3.bucket),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return fmt.Errorf("aws error: error uploading file to aws %v", err)
	}

	return nil
}

func (s *Storage) GetResource(ctx context.Context, id string) (*models.Resource, error) {
	resourceUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("db error: invalid resource id")
	}

	resourceDB, err := s.queries.GetResourceByID(ctx, resourceUUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("db error: getting resource from database")
	}

	resourceURL, err := url.Parse(resourceDB.S3Url)
	if err != nil {
		return nil, fmt.Errorf("somethign deeply is wrong here, s3 url in not a valid url: %v", err)
	}

	return &models.Resource{
		ID:          &resourceDB.ID,
		CourseID:    resourceDB.CourseID,
		Title:       resourceDB.Title,
		Description: resourceDB.Description,
		FileExt:     resourceDB.FileExt,
		Url:         resourceURL,
		Tags:        resourceDB.Tags,
		Created_by:  resourceDB.CreatedBy,
	}, nil
}

func (s *Storage) GetResourceDownloadURL(ctx context.Context, resource *models.Resource) (string, error) {

	presignClient := s3.NewPresignClient(s.s3.Client)
	downloadReq, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket:                     aws.String(s.s3.bucket),
		Key:                        aws.String(s.fileS3Key(resource.CourseID.String(), resource.Title, resource.FileExt)),
		ResponseContentDisposition: aws.String(fmt.Sprintf("%s%s", resource.Title, resource.FileExt)),
	}, s3.WithPresignExpires(15*time.Minute))
	if err != nil {
		return "", fmt.Errorf("s3 error: error generating download link %v", err)
	}

	return downloadReq.URL, nil
}

func (s *Storage) fileS3Key(courseId, fileName, ext string) string {
	return fmt.Sprintf("%s/%s%s", courseId, fileName, ext)

}

func (s *Storage) fileS3URL(key string) string {
	urlSchema := "https://%s.s3-%s.amazonaws.com/%s"

	return fmt.Sprintf(urlSchema, s.s3.bucket, s.s3.region, key)
}
