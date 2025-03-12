package storage

import (
	"database/sql"

	"github.com/Modalessi/iau_resources/database"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Storage struct {
	s3      *S3Config
	db      *sql.DB
	queries *database.Queries
}

type S3Config struct {
	Client *s3.Client
	bucket string
	region string
}

func NewS3Config(client *s3.Client, bucket, region string) *S3Config {
	return &S3Config{
		Client: client,
		bucket: bucket,
		region: region,
	}
}

func NewStorage(db *sql.DB, queries *database.Queries, s3config *S3Config) *Storage {
	return &Storage{
		s3:      s3config,
		db:      db,
		queries: queries,
	}
}
