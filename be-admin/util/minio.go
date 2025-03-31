package config

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinIO Configuration Struct
type MinioConfig struct {
	Client      *minio.Client
	BucketName  string
	FileUrl     string
	VideoConfig *VideoConfig
}

// VideoConversionConfig struct for settings
type VideoConfig struct {
	OutputFormat string
	SegmentTime  int
	VideoQuality string
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

	return &MinioConfig{Client: client, BucketName: bucketName, FileUrl: fileUrl, VideoConfig: &VideoConfig{
		OutputFormat: "libx264", // Default to H.264
		SegmentTime:  5,         // 5 seconds per segment
		VideoQuality: "medium",  // Default quality
	}}
}

// SaveImage uploads an image file to MinIO
func (m *MinioConfig) SaveImage(fileHeader *multipart.FileHeader, folderPath string) (string, error) {
	if err := ValidateImageFile(fileHeader); err != nil {
		return "", err
	}
	return m.UploadFile(fileHeader, folderPath)
}

// SaveVideo uploads a video file to MinIO
func (m *MinioConfig) SaveVideo(fileHeader *multipart.FileHeader, folderPath string, id int64) (string, error) {
	if err := ValidateVideoFile(fileHeader); err != nil {
		return "", err
	}
	return m.ConvertToHLSAndSave(fileHeader, folderPath, id)
}

// SaveSubtitle uploads a subtitle file to MinIO
func (m *MinioConfig) SaveSubtitle(fileHeader *multipart.FileHeader, folderPath string) (string, error) {
	if err := ValidateSubtitleFile(fileHeader); err != nil {
		return "", err
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

func (m *MinioConfig) UploadFileByte(objectName string, fileData []byte) error {
	// Set MIME type based on file extension
	contentType := "application/octet-stream"
	if strings.HasSuffix(objectName, ".m3u8") {
		contentType = "application/vnd.apple.mpegurl"
	} else if strings.HasSuffix(objectName, ".ts") {
		contentType = "video/mp2t"
	}

	reader := bytes.NewReader(fileData)
	if _, err := m.Client.PutObject(context.Background(), m.BucketName, objectName, reader, int64(len(fileData)),
		minio.PutObjectOptions{ContentType: contentType}); err != nil {
		return fmt.Errorf("failed to upload %s: %v", objectName, err)
	}
	return nil
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

// ConvertToHLS converts video and uploads it to MinIO
func (m *MinioConfig) ConvertToHLSAndSave(fileHeader *multipart.FileHeader, folderPath string, dramaId int64) (string, error) {
	outputName := MakeUniqueIDWithTime()
	dramaIdStr := strconv.FormatInt(dramaId, 10)

	// Create temporary directory
	tempDir := fmt.Sprintf("temp_%s", outputName)
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up temp directory after we're done

	// Save uploaded file to temp directory
	tempInputPath := filepath.Join(tempDir, fileHeader.Filename)
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %v", err)
	}
	defer file.Close()

	// Create the temporary file
	tempFile, err := os.Create(tempInputPath)
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %v", err)
	}
	defer tempFile.Close()

	// Copy the uploaded file to temp file
	if _, err := tempFile.ReadFrom(file); err != nil {
		return "", fmt.Errorf("failed to copy file: %v", err)
	}

	// Generate unique output names
	m3u8FileName := outputName + ".m3u8"
	m3u8File := filepath.Join(tempDir, m3u8FileName)
	tsPattern := filepath.Join(tempDir, outputName+"_%03d.ts")

	// Build FFmpeg command
	args := []string{
		"-i", tempInputPath, // Input video file
		"-c:v", m.VideoConfig.OutputFormat, // Video codec (e.g., libx264, libx265)
		"-c:a", "aac", // Audio codec (AAC for compatibility)
		"-preset", "veryfast", // Faster encoding (adjustable: ultrafast, superfast, veryfast, fast, etc.)
		"-tune", "zerolatency", // Optimize for low-latency streaming
		"-hls_time", fmt.Sprintf("%d", m.VideoConfig.SegmentTime), // Segment duration in seconds
		"-hls_playlist_type", "vod", // Set HLS playlist type (VOD for full playlist storage)
		"-hls_segment_filename", tsPattern, // Naming pattern for .ts segments
		"-hls_flags", "independent_segments", // Allow independent segment playback
		"-hls_list_size", "0", // Keep all segments in the playlist (important for VOD)
		"-movflags", "+faststart", // Enable progressive download (better for streaming)
		"-max_muxing_queue_size", "1024", // Avoid queue overflow issues
		"-thread_queue_size", "512", // Increase buffer size for input queue
		"-threads", fmt.Sprintf("%d", runtime.NumCPU()/2), // Set thread count (using half of available CPU cores)
		"-f", "hls", // Output format (HLS)
		m3u8File, // Output M3U8 playlist file
	}

	// Set quality settings
	switch m.VideoConfig.VideoQuality {
	case "high":
		args = append(args, "-crf", "17", "-preset", "slow")
	case "medium":
		args = append(args, "-crf", "23", "-preset", "medium")
	case "low":
		args = append(args, "-crf", "28", "-preset", "fast")
	}

	// Create FFmpeg command
	cmd := exec.Command("ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run FFmpeg
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("FFmpeg failed: %v", err)
	}

	// Read .m3u8 file content
	m3u8Content, err := os.ReadFile(m3u8File)
	if err != nil {
		return "", fmt.Errorf("failed to read .m3u8 file: %v", err)
	}

	// Extract .ts filenames
	tsRegex := regexp.MustCompile(fmt.Sprintf(`%s_\d{3}\.ts`, outputName))
	tsMatches := tsRegex.FindAllString(string(m3u8Content), -1)

	// Upload .m3u8 file
	result := fmt.Sprintf("%s/%s", dramaIdStr, m3u8FileName)
	m3u8ObjectName := fmt.Sprintf("%s%s", folderPath, result)
	if err := m.UploadFileByte(m3u8ObjectName, m3u8Content); err != nil {
		return "", err
	}

	// Upload .ts files
	errChan := make(chan error, len(tsMatches))
	var wg sync.WaitGroup

	for _, tsFile := range tsMatches {
		wg.Add(1)
		go func(tsFile string) {
			defer wg.Done()

			tsData, err := os.ReadFile(filepath.Join(tempDir, tsFile))
			if err != nil {
				errChan <- fmt.Errorf("failed to read %s: %v", tsFile, err)
				return
			}

			tsObjectName := fmt.Sprintf("%s%s/%s", folderPath, dramaIdStr, tsFile)
			if err := m.UploadFileByte(tsObjectName, tsData); err != nil {
				errChan <- err
				return
			}
		}(tsFile)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		return "", err
	}

	return result, nil
}
