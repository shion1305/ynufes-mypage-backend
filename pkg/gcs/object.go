package gcs

import (
	"cloud.google.com/go/storage"
	"context"
	"io"
	"time"
)

type (
	ObjectRef struct {
		objName string
		bucket  *storage.BucketHandle
	}
	SignedObjectLink string
)

func (r ObjectRef) Name() string {
	return r.objName
}

func (r ObjectRef) Download(ctx context.Context) ([]byte, error) {
	reader, err := r.bucket.Object(r.objName).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r ObjectRef) IssueLink() (SignedObjectLink, error) {
	url, err := r.bucket.SignedURL(r.objName, &storage.SignedURLOptions{
		Method:  "GET",
		Expires: time.Now().Add(time.Minute * 1),
		Scheme:  storage.SigningSchemeV4,
	})
	if err != nil {
		return "", err
	}
	return SignedObjectLink(url), nil
}
