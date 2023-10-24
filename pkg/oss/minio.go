package oss

import (
	"context"
	"fmt"
	"github.com/luyasr/simple-blog/config"
	"github.com/luyasr/simple-blog/pkg/logger"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog"
	"net/url"
)

var (
	_ Service = (*Minio)(nil)
)

type Minio struct {
	core       *minio.Core
	bucketName string
	log        zerolog.Logger
}

func NewMinio() *Minio {
	return &Minio{
		core:       config.C.Minio.NewCore(),
		bucketName: config.C.Minio.BucketName,
		log:        logger.NewConsoleLog(),
	}
}

func (m *Minio) MakeBucket(ctx context.Context) error {
	err := m.core.MakeBucket(ctx, m.bucketName, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		exists, errBucketExists := m.core.BucketExists(ctx, m.bucketName)
		if errBucketExists == nil && exists {
			m.log.Info().Msg(fmt.Sprintf("We already own %s", m.bucketName))
			return errBucketExists
		} else {
			m.log.Error().Err(err).Send()
			return err
		}
	}

	m.log.Info().Msg(fmt.Sprintf("Successfully created %s", m.bucketName))
	return nil
}

func (m *Minio) GetMultipartUploadId(ctx context.Context, req *GetMultipartUploadIdRequest) (string, error) {
	uploadID, err := m.core.NewMultipartUpload(ctx, m.bucketName, req.ObjectName, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}

	return uploadID, nil
}

func (m *Minio) GetPresignedURL(ctx context.Context, req *GetPresignedURLRequest) (*url.URL, error) {
	u, err := m.core.Presign(ctx, "PUT", m.bucketName, req.ObjectName, req.Expires, req.params())
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (m *Minio) CompleteMultipartUpload(ctx context.Context, req *CompleteMultipartUploadRequest) (*minio.UploadInfo, error) {
	uploadInfo, err := m.core.CompleteMultipartUpload(ctx, m.bucketName, req.ObjectName, req.UploadID, req.CompleteParts, minio.PutObjectOptions{})
	if err != nil {
		return nil, err
	}

	return &uploadInfo, nil
}
