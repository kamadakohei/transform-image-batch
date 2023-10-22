package gcs

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func DownloadFile(ctx context.Context, client *storage.Client, w io.Writer, bucket, object string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	destFileName := "./gcs/tmp/download/" + filepath.Base(object)
	f, err := os.Create(destFileName)
	if err != nil {
		return "", fmt.Errorf("os.Create: %w", err)
	}

	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return "", fmt.Errorf("Object(%q).NewReader: %w", object, err)
	}
	defer rc.Close()

	if _, err := io.Copy(f, rc); err != nil {
		return "", fmt.Errorf("io.Copy: %w", err)
	}

	if err = f.Close(); err != nil {
		return "", fmt.Errorf("f.Close: %w", err)
	}

	fmt.Fprintf(w, "Blob %v downloaded to local file %v\n", object, destFileName)

	return destFileName, nil

}

func UploadFile(w io.Writer, bucket, objectFilePath string) error {
	uploadFilePath := "upload/" + filepath.Base(objectFilePath)

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()
	currentDir, err := os.Getwd()
	fmt.Printf(currentDir)
	f, err := os.Open(objectFilePath)
	if err != nil {
		return fmt.Errorf("os.Open: %w", err)
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	o := client.Bucket(bucket).Object(uploadFilePath)

	// Check if the object exists before overwriting.
	o = o.If(storage.Conditions{DoesNotExist: true})

	wc := o.NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %w", err)
	}
	fmt.Fprintf(w, "Blob %v uploaded.\n", uploadFilePath)
	return nil
}
