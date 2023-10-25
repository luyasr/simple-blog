package oss

import (
	"context"
	"github.com/minio/minio-go/v7"
	"net/url"
	"time"
)

type Service interface {
	MakeBucket(context.Context) error
	MultipartUpload(context.Context, *MultipartUploadRequest) (*MultipartUploadResponse, error)
	CompleteMultipartUpload(context.Context, *CompleteMultipartUploadRequest) (*minio.UploadInfo, error)
}

type MultipartUploadRequest struct {
	ObjectName string `json:"object_name"`
	SumSHA256  string `json:"sum_SHA256"`
	Size       int64  `json:"size"`
	PartSize   int64
	Expires    time.Duration
}

func NewMultipartUploadRequest() *MultipartUploadRequest {
	return &MultipartUploadRequest{
		PartSize: 5 * 1024 * 1024,
		Expires:  time.Second * 3600,
	}
}

type MultipartUploadResponse struct {
	UploadId   string     `json:"upload_id"`
	ObjectName string     `json:"object_name"`
	PartSize   int64      `json:"part_size"`
	PartInfo   []PartInfo `json:"part_info"`
}

type PartInfo struct {
	PartNumber int      `json:"part_number"`
	PresignURL *url.URL `json:"presign_url"`
}

type CompleteMultipartUploadRequest struct {
	UploadID      string `json:"upload_id"`
	ObjectName    string `json:"object_name"`
	CompleteParts []minio.CompletePart
}

func NewCompleteMultipartUploadRequest() *CompleteMultipartUploadRequest {
	return &CompleteMultipartUploadRequest{}
}
