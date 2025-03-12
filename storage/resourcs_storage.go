package storage

import (
	"context"
	"fmt"
	"io"
	"net/url"

	"github.com/Modalessi/iau_resources/database"
	"github.com/Modalessi/iau_resources/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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

func (s *Storage) fileS3Key(courseId, fileName, ext string) string {
	return fmt.Sprintf("%s/%s%s", courseId, fileName, ext)

}

func (s *Storage) fileS3URL(key string) string {
	urlSchema := "https://%s.s3-%s.amazonaws.com/%s"

	return fmt.Sprintf(urlSchema, s.s3.bucket, s.s3.region, key)
}
