package gcs

import "cloud.google.com/go/storage"

type BucketRef struct {
	bucket *storage.BucketHandle
}

func (r BucketRef) Folder(name string) *FolderRef {
	return &FolderRef{
		basePath: name,
		bucket:   nil,
	}
}
