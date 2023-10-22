package gcs

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
)

func Client() {
	client, err := storage.NewClient(context.TODO(), option.WithEndpoint("http://0.0.0.0:4443/storage/v1/"))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	const (
		bucketName = "some-bucket"
	)
	buckets, err := list(client, bucketName)
	if err != nil {
		log.Fatalf("failed to list: %v", err)
	}
	fmt.Printf("buckets: %+v\n", buckets)
}

func list(client *storage.Client, bucketName string) ([]string, error) {
	var objects []string
	it := client.Bucket(bucketName).Objects(context.Background(), &storage.Query{})
	for {
		oattrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		objects = append(objects, oattrs.Name)
	}
	return objects, nil
}
