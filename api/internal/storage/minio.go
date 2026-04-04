package storage

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	Client *minio.Client
	Bucket string
}

func NewMinioClient(endpoint, accessKey, secretKey, bucket string, useSSL bool) (*MinioClient, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}
	return &MinioClient{Client: client, Bucket: bucket}, nil
}

func (m *MinioClient) UploadFile(ctx context.Context, objectName string, file io.Reader, size int64, contentType string) error {
	_, err := m.Client.PutObject(ctx, m.Bucket, objectName, file, size, minio.PutObjectOptions{ContentType: contentType})
	return err
}
