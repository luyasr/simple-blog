package oss

import (
	"context"
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
	req.ObjectName = "test/c.jpeg"

	uploadId, err := m.GetMultipartUploadId(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uploadId)
}

func TestMinio_GetPresignedURL(t *testing.T) {
	req := NewGetPresignedURLRequest()
	req.UploadID = "YTIzMjQ3YTAtNjEwZi00YzgyLWFjNjktNmEyNGFkZmRhNmM1Ljc4MjZjM2ZjLWEzM2QtNGIxNS04ZjhhLWE5M2FlMzc5NWU1Yw"
	req.ObjectName = "test/c.jpeg"
	req.PartNumber = 1
	url, err := m.GetPresignedURL(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
}

func TestMinio_CompleteMultipartUpload(t *testing.T) {
	req := NewCompleteMultipartUploadRequest()
	req.ObjectName = "test/c.jpeg"
	req.UploadID = "YTIzMjQ3YTAtNjEwZi00YzgyLWFjNjktNmEyNGFkZmRhNmM1Ljc4MjZjM2ZjLWEzM2QtNGIxNS04ZjhhLWE5M2FlMzc5NWU1Yw"
	uploadInfo, err := m.CompleteMultipartUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uploadInfo)
}
