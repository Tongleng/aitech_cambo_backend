package configs

import (
	"context"
	"io"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var StorageClient *storage.Client
var BucketName = "your-bucket-name"

func InitStorage() error {
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile("service-account.json"))
	if err != nil {
		return err
	}

	StorageClient = client
	return nil
}

func UploadFile(file multipart.File, filename string) (string, error) {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := StorageClient.Bucket(BucketName).Object(filename).NewWriter(ctx)

	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	url := "https://storage.googleapis.com/" + BucketName + "/" + filename

	return url, nil
}
