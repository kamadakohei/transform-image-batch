package gcs

import (
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/option"
	"log"
)

func InitGCSClient() (context.Context, *storage.Client, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	return ctx, client, err
}

func InitGCSTestClient() (context.Context, *storage.Client, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithEndpoint("http://0.0.0.0:4443/storage/v1/"))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	return ctx, client, err
}
