package oss

import (
	"context"
	"fmt"
	"github.com/luyasr/simple-blog/config"
	"github.com/luyasr/simple-blog/pkg/logger"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"math"
	"net/url"
	"strconv"
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

func (m *Minio) MultipartUpload(ctx context.Context, req *MultipartUploadRequest) (*MultipartUploadResponse, error) {
	var g errgroup.Group
	var partInfos []PartInfo

	uploadId, err := m.core.NewMultipartUpload(ctx, m.bucketName, req.ObjectName, minio.PutObjectOptions{})
	if err != nil {
		return nil, err
	}

	partNumbers := sharding(req.Size, req.PartSize)
	for i := 1; i <= partNumbers; i++ {
		partNumber := i
		g.Go(func() error {
			urlValues := make(url.Values)
			urlValues.Set("uploadId", uploadId)
			urlValues.Set("partNumber", strconv.Itoa(partNumber))

			partUrl, err := m.core.Presign(ctx, "PUT", m.bucketName, req.ObjectName, req.Expires, urlValues)
			if err != nil {
				return err
			}
			partInfos = append(partInfos, PartInfo{PartNumber: partNumber, PresignURL: partUrl})

			return nil
		})
	}

	if err = g.Wait(); err != nil {
		return nil, err
	}

	return &MultipartUploadResponse{
		UploadId: uploadId,
		PartSize: req.PartSize,
		PartInfo: partInfos,
	}, nil
}

func (m *Minio) CompleteMultipartUpload(ctx context.Context, req *CompleteMultipartUploadRequest) (*minio.UploadInfo, error) {
	result, err := m.core.ListObjectParts(ctx, m.bucketName, req.ObjectName, req.UploadID, 0, 10000)
	if err != nil {
		return nil, err
	}

	fmt.Println(result.ObjectParts)
	for _, part := range result.ObjectParts {
		req.CompleteParts = append(req.CompleteParts, minio.CompletePart{PartNumber: part.PartNumber, ETag: part.ETag})
	}

	uploadInfo, err := m.core.CompleteMultipartUpload(ctx, m.bucketName, req.ObjectName, req.UploadID, req.CompleteParts, minio.PutObjectOptions{})
	if err != nil {
		return nil, err
	}

	return &uploadInfo, nil
}

func sharding(objectSize int64, partSize int64) int {
	var shard float64
	if objectSize > partSize {
		shard = math.Ceil(float64(objectSize) / float64(partSize))
	} else {
		shard = 1
	}

	return int(shard)
}
