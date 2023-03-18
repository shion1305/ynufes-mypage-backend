package gcs

import (
	"cloud.google.com/go/storage"
	"time"
)

type (
	ObjectRef struct {
		objName string
		bucket  *storage.BucketHandle
	}
	SignedObjectLink string
)

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
