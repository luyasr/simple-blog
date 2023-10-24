package oss

import (
	"context"
	"github.com/minio/minio-go/v7"
	"net/url"
	"strconv"
	"time"
)

type Service interface {
	MakeBucket(context.Context) error
	GetMultipartUploadId(context.Context, *GetMultipartUploadIdRequest) (string, error)
	GetPresignedURL(context.Context, *GetPresignedURLRequest) (*url.URL, error)
	CompleteMultipartUpload(context.Context, *CompleteMultipartUploadRequest) (*minio.UploadInfo, error)
}

type GetMultipartUploadIdRequest struct {
	ObjectName string `json:"object_name"`
}

func NewGetMultipartUploadIdRequest() *GetMultipartUploadIdRequest {
	return &GetMultipartUploadIdRequest{}
}

type GetPresignedURLRequest struct {
	UploadID   string        `json:"upload_id"`
	PartNumber int           `json:"part_number"`
	ObjectName string        `json:"object_name"`
	Expires    time.Duration `json:"expires"`
}

func (p *GetPresignedURLRequest) params() url.Values {
	return url.Values{
		"uploadId":   []string{p.UploadID},
		"partNumber": []string{strconv.Itoa(p.PartNumber)},
	}
}

func NewGetPresignedURLRequest() *GetPresignedURLRequest {
	return &GetPresignedURLRequest{
		Expires: time.Second * 3600,
	}
}

type CompleteMultipartUploadRequest struct {
	UploadID      string               `json:"upload_id"`
	ObjectName    string               `json:"object_name"`
	CompleteParts []minio.CompletePart `json:"complete_parts"`
}

func NewCompleteMultipartUploadRequest() *CompleteMultipartUploadRequest {
	return &CompleteMultipartUploadRequest{}
}
