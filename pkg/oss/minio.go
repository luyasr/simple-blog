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
	"mime"
	"net/url"
	"path"
	"sort"
	"strconv"
)

var (
	_ Service = (*Minio)(nil)
)

const (
	DefaultRegion = "us-east-1"
	MaxParts      = 10000
)

type Minio struct {
	core       *minio.Core
	bucketName string
	log        zerolog.Logger
}

func NewMinio() *Minio {
	return &Minio{
		core:       config.C.Minio.GetCore(),
		bucketName: config.C.Minio.BucketName,
		log:        logger.NewConsoleLog(),
	}
}

func (m *Minio) MakeBucket(ctx context.Context) error {
	err := m.core.MakeBucket(ctx, m.bucketName, minio.MakeBucketOptions{Region: DefaultRegion})
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

func (m *Minio) getLastUploadID(ctx context.Context, objectName string) (string, error) {
	var lastUploadID string
	incompleteUploads := m.core.ListIncompleteUploads(ctx, m.bucketName, objectName, true)

	for upload := range incompleteUploads {
		if upload.Err != nil {
			return "", upload.Err
		}
		lastUploadID = upload.UploadID
	}
	return lastUploadID, nil
}

func (m *Minio) MultipartUpload(ctx context.Context, req *MultipartUploadRequest) (*MultipartUploadResponse, error) {
	var g errgroup.Group
	partInfoCh := make(chan PartInfo)

	// 对象是否已经存在
	_, err := m.core.StatObject(ctx, m.bucketName, req.ObjectName, minio.StatObjectOptions{})
	if err == nil {
		return &MultipartUploadResponse{Uploaded: true}, nil
	}

	// minio断点续传, 获取上次上传的uploadId
	lastUploadId, err := m.getLastUploadID(ctx, req.ObjectName)
	if err != nil {
		return nil, err
	}

	// 获取对象分片信息
	result, err := m.core.ListObjectParts(ctx, m.bucketName, req.ObjectName, lastUploadId, 0, MaxParts)
	if err != nil {
		return nil, err
	}

	// 获取已经上传的分片
	var uploadId string
	objParts := make(map[int]struct{})
	if len(result.ObjectParts) > 0 {
		for _, part := range result.ObjectParts {
			objParts[part.PartNumber] = struct{}{}
		}
		uploadId = lastUploadId
	}

	// 如果上次uploadId存在已经上传的分片, 则延用上次uploadId
	if uploadId == "" {
		// 获取对象uploadId
		uploadId, err = m.core.NewMultipartUpload(ctx, m.bucketName, req.ObjectName, minio.PutObjectOptions{
			ContentType: m.ContentType(req.ObjectName),
		})

		if err != nil {
			return nil, err
		}
	}

	// 分片预签名
	partNumbers := m.Sharding(req.Size, req.PartSize)
	for i := 1; i <= partNumbers; i++ {
		partNumber := i
		if _, ok := objParts[partNumber]; ok {
			continue
		}
		g.Go(func() error {
			urlValues := make(url.Values)
			urlValues.Set("uploadId", uploadId)
			urlValues.Set("partNumber", strconv.Itoa(partNumber))

			partUrl, err := m.core.Presign(ctx, "PUT", m.bucketName, req.ObjectName, req.Expires, urlValues)
			if err != nil {
				return err
			}

			partInfoCh <- PartInfo{PartNumber: partNumber, PresignURL: partUrl}

			return nil
		})
	}

	// 等待所有协程执行完毕后关闭通道
	go func() {
		_ = g.Wait()
		close(partInfoCh)
	}()

	// 从通道中获取分片信息
	var partInfos []PartInfo
	for partInfo := range partInfoCh {
		partInfos = append(partInfos, partInfo)
	}

	// 获取协程中返回的报错
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
	result, err := m.core.ListObjectParts(ctx, m.bucketName, req.ObjectName, req.UploadId, 0, MaxParts)
	if err != nil {
		return nil, err
	}

	for _, part := range result.ObjectParts {
		req.CompleteParts = append(req.CompleteParts, minio.CompletePart{PartNumber: part.PartNumber, ETag: part.ETag})
	}

	// 合并前需要对分片进行排序 https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html
	sort.Slice(req.CompleteParts, func(i, j int) bool {
		return req.CompleteParts[i].PartNumber < req.CompleteParts[j].PartNumber
	})

	uploadInfo, err := m.core.CompleteMultipartUpload(ctx, m.bucketName, req.ObjectName, req.UploadId, req.CompleteParts, minio.PutObjectOptions{
		ContentType: m.ContentType(req.ObjectName),
	})
	if err != nil {
		return nil, err
	}

	return &uploadInfo, nil
}

func (m *Minio) AbortMultipartUpload(ctx context.Context, req *AbortMultipartUploadRequest) error {
	return m.core.AbortMultipartUpload(ctx, m.bucketName, req.ObjectName, req.UploadId)
}

func (m *Minio) ContentType(objectName string) string {
	ext := path.Ext(objectName)
	contentType := mime.TypeByExtension(ext)

	if contentType == "" {
		contentType = "application/octet-stream"
	}
	return contentType
}

func (m *Minio) Sharding(objectSize int64, partSize int64) int {
	var shard float64
	if objectSize > partSize {
		shard = math.Ceil(float64(objectSize) / float64(partSize))
	} else {
		shard = 1
	}

	return int(shard)
}
