package gcs

import (
	"cloud.google.com/go/storage"
	"context"
)

type Client struct {
	client *storage.Client
}

func NewClient(ctx context.Context) (*Client, error) {
	client, err := storage.NewClient(ctx)

	if err != nil {
		return nil, err
	}
	return &Client{client: client}, nil
}

func (c Client) Bucket(name string) *BucketRef {
	return &BucketRef{
		bucket: c.client.Bucket(name),
	}
}
