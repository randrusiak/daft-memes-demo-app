package main

import (
	"context"
	"errors"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GCSUploadFile(file io.Reader, filename string) (string, error) {
	bucket := os.Getenv("GCS_BUCKET_NAME")
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", errors.New("failed to create client")
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Minute*20)
	defer cancel()

	wc := client.Bucket(bucket).Object(filename).NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		return "", errors.New("failed to copy file")
	}
	if err := wc.Close(); err != nil {
		return "", errors.New("failed to save file")
	}
	return "https://storage.googleapis.com/" + bucket + "/" + filename, nil
}
