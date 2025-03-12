package main

import (
	"context"
	"database/sql"
	"os"

	"github.com/Modalessi/iau_resources/database"
	fanar "github.com/Modalessi/iau_resources/fanar_api"
	"github.com/Modalessi/iau_resources/storage"
	"github.com/Modalessi/iau_resources/utils"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load()

	DB_URL := os.Getenv("DB_URL")
	utils.Assert(DB_URL != "", "could not read DB_URL from .env file, please check that it exist")

	JWT_SECRET := os.Getenv("JWT_SECRET")
	utils.Assert(JWT_SECRET != "", "could not read JWT_SECRET from .env file, please check that it exist")

	ADMIN_EMAIL := os.Getenv("ADMIN_EMAIL")
	utils.Assert(ADMIN_EMAIL != "", "could not read ADMIN_EMAIL from .env file, please check that it exist")

	db, err := sql.Open("postgres", DB_URL)
	utils.ErrorAssert(err, "error connecting to database")

	quries := database.New(db)

	awsConfig, err := config.LoadDefaultConfig(context.TODO())
	utils.ErrorAssert(err, "error creating aws config")

	AWS_FILES_BUCKET := os.Getenv("AWS_FILES_BUCKET")
	utils.Assert(AWS_FILES_BUCKET != "", "somethign went wrong when reading 'AWS_FILES_BUCKET' env variable")

	AWS_REGION := os.Getenv("AWS_REGION")
	utils.Assert(AWS_REGION != "", "somethign went wrong when reading 'AWS_REGION' env variable")

	s3client := s3.NewFromConfig(awsConfig)

	storageS3Config := storage.NewS3Config(s3client, AWS_FILES_BUCKET, AWS_REGION)

	fanarStorage := storage.NewStorage(db, quries, storageS3Config)
	fanarServer := fanar.NewFanarServer(":4000", JWT_SECRET, ADMIN_EMAIL, fanarStorage)

	err = fanarServer.Start()
	if err != nil {
		panic("error starting fanar server")
	}
}
