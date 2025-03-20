package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"strconv"
	"sync"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinIO Configuration Struct
type MinioConfig struct {
	Client     *minio.Client
	BucketName string
	FileUrl    string
}

// Initialize MinIO Client
func NewMinioClient(endpoint, accessKey, secretKey, bucketName string, useSSL bool, fileUrl string) *MinioConfig {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("MinIO connection failed: %v", err)
	}

	// Check if bucket exists, create if not
	exists, err := client.BucketExists(context.Background(), bucketName)
	if err != nil {
		log.Fatalf("Bucket check failed: %v", err)
	}
	if !exists {
		err = client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("Could not create bucket: %v", err)
		}
	}

	return &MinioConfig{Client: client, BucketName: bucketName, FileUrl: fileUrl}
}

// SaveImage uploads an image file to MinIO
func (m *MinioConfig) SaveImage(fileHeader *multipart.FileHeader, folderPath string) (string, error) {
	if fileHeader.Size > MaxImageSize {
		return "", errors.New("file size exceeds the maximum allowed limit of 1MB")
	}

	isImage, err := IsImage(fileHeader)
	if err != nil {
		return "", err
	}

	if !isImage {
		return "", errors.New("file is not an image")
	}

	return m.UploadFile(fileHeader, folderPath)
}

// SaveVideo uploads a video file to MinIO
func (m *MinioConfig) SaveVideo(fileHeader *multipart.FileHeader, folderPath string) (string, error) {
	if fileHeader.Size > MaxVideoSize {
		return "", errors.New("file size exceeds the maximum allowed limit of 1GB")
	}

	isVideo, err := IsVideo(fileHeader)
	if err != nil {
		return "", err
	}

	if !isVideo {
		return "", errors.New("file is not a video")
	}

	return m.UploadFile(fileHeader, folderPath)
}

func (m *MinioConfig) UploadFile(fileHeader *multipart.FileHeader, folderPath string) (string, error) {
	ctx := context.Background()
	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	name := fmt.Sprintf("%s%s", MakeUniqueIDWithTime(), GetFileExtension(fileHeader))
	objectName := fmt.Sprintf("%s%s", folderPath, name)

	// Upload file
	_, err = m.Client.PutObject(ctx, m.BucketName, objectName, file, fileHeader.Size, minio.PutObjectOptions{
		ContentType: fileHeader.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", err
	}

	return name, nil
}

func (m *MinioConfig) GetFileUrl(folderPath string, fileName string) string {
	return fmt.Sprintf("%s/%s%s%s", m.FileUrl, m.BucketName, folderPath, fileName)
}

func (m *MinioConfig) SaveImages(fileHeaders []*multipart.FileHeader, folderPath string) ([]string, error) {
	ctx := context.Background()
	names := make([]string, len(fileHeaders))
	errChan := make(chan error, len(fileHeaders))
	var wg sync.WaitGroup

	for i, fileHeader := range fileHeaders {
		wg.Add(1)
		go func(index int, fileHeader *multipart.FileHeader) {
			defer wg.Done()

			if fileHeader.Size > MaxImageSize {
				errChan <- errors.New("file size exceeds the maximum allowed limit of 1MB")
				return
			}

			isImage, err := IsImage(fileHeader)
			if err != nil || !isImage {
				errChan <- errors.New("file is not an image")
				return
			}

			// Open the file
			file, err := fileHeader.Open()
			if err != nil {
				errChan <- err
				return
			}
			defer file.Close()

			name := fmt.Sprintf("%s%s%s", MakeUniqueIDWithTime(), strconv.Itoa(i), GetFileExtension(fileHeader))
			objectName := fmt.Sprintf("%s%s", folderPath, name)

			// Upload file
			_, err = m.Client.PutObject(ctx, m.BucketName, objectName, file, fileHeader.Size, minio.PutObjectOptions{
				ContentType: fileHeader.Header.Get("Content-Type"),
			})

			if err != nil {
				errChan <- err
				return
			}

			names[index] = name
		}(i, fileHeader)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		return nil, err
	}

	return names, nil
}
