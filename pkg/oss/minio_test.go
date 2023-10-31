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
	req.ObjectName = "avatar/uTools-4.1.0.exe"
	req.Size = 66883608

	response, err := m.MultipartUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestMinio_CompleteMultipartUpload(t *testing.T) {
	req := NewCompleteMultipartUploadRequest()
	req.UploadId = "MzI3YTVkNmUtNTExMy00YzRkLWJhNGEtYmEyNjJiZjRhZmE2LjllYTMwZDYzLTFmZTYtNDRhMi05ZWYzLTI1MDEyNDkwODY3Yg"
	req.ObjectName = "avatar/uTools-4.1.0.exe"

	info, err := m.CompleteMultipartUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestMinio_AbortMultipartUpload(t *testing.T) {
	req := NewAbortMultipartUploadRequest()
	req.UploadId = "MzI3YTVkNmUtNTExMy00YzRkLWJhNGEtYmEyNjJiZjRhZmE2LmYwZTk1MjQ4LTI3MTAtNGZjNi04MTBhLTQyNTFkYWFhMjRlOQ"
	req.ObjectName = "avatar/uTools-4.1.0.exe"

	err := m.AbortMultipartUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}