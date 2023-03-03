package config

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/joho/godotenv"
)

type ClientUploader struct {
	Cl         *storage.Client
	ProjectID  string
	BucketName string
	UploadPath string
}

func SetupCloudStorage() *ClientUploader {

	var uploader *ClientUploader
	// .ENV VAR
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load .env file")
	}

	bucketName := os.Getenv("BUCKET_NAME")
	gcsKeyPath := os.Getenv("GCS_FILE_PATH")
	projectID := os.Getenv("PROJECT_ID")

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", gcsKeyPath) // FILL IN WITH YOUR FILE PATH

	client, err := storage.NewClient(context.Background())

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	uploader = &ClientUploader{
		Cl:         client,
		BucketName: bucketName,
		ProjectID:  projectID,
		UploadPath: "project1/upload-file/",
	}
	return uploader
}
