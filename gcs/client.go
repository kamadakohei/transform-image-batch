package gcs

import (
	"cloud.google.com/go/storage"
	"context"
)

func InitGCSClient() (context.Context, *storage.Client, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	return ctx, client, err
}
