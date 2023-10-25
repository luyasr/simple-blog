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

func TestMinio_MultipartUpload(t *testing.T) {
	req := NewMultipartUploadRequest()
	req.ObjectName = "avatar/image.jpg"
	req.Size = 4 * 1024 * 1024

	response, err := m.MultipartUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestMinio_CompleteMultipartUpload(t *testing.T) {
	req := NewCompleteMultipartUploadRequest()
	req.UploadID = "NjAyYmJlYzctYjI0MC00OWIwLThjOTAtMGZjN2QxODhiYzJhLmRhYzRlZmI4LTFjMzYtNGYzMi05YTQ0LTZiYmZlMzRiNmU4Mg"
	req.ObjectName = "avatar/image.jpg"

	info, err := m.CompleteMultipartUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}
