package oss

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"testing"
)

var (
	m   Service
	ctx = context.Background()
)

func init() {
	m = NewMinio()
}

func TestMinio_MakeBucket(t *testing.T) {
	err := m.MakeBucket(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMinio_GetMultipartUploadId(t *testing.T) {
	req := NewGetMultipartUploadIdRequest()
	req.ObjectName = "test/头像.jpg"

	uploadId, err := m.GetMultipartUploadId(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uploadId)
}

func TestMinio_GetPresignedURL(t *testing.T) {
	req := NewGetPresignedURLRequest()
	req.UploadID = "NjAyYmJlYzctYjI0MC00OWIwLThjOTAtMGZjN2QxODhiYzJhLmQxYzIyNzdjLWI4MzYtNGRlZi1iN2Y4LWY2NTQ5ODMyNjIwZg"
	req.ObjectName = "test/头像.jpg"
	req.PartNumber = 1
	url, err := m.GetPresignedURL(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
}

func TestMinio_CompleteMultipartUpload(t *testing.T) {
	req := NewCompleteMultipartUploadRequest()
	req.ObjectName = "test/头像.jpg"
	req.CompleteParts = []minio.CompletePart{{PartNumber: 1}}
	fmt.Println(req)
	req.UploadID = "NjAyYmJlYzctYjI0MC00OWIwLThjOTAtMGZjN2QxODhiYzJhLmQxYzIyNzdjLWI4MzYtNGRlZi1iN2Y4LWY2NTQ5ODMyNjIwZg"

	uploadInfo, err := m.CompleteMultipartUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uploadInfo)
}
