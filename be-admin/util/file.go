package config

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"
)

const MaxImageSize = 1 << 20 // 1MB
const MaxVideoSize = 1 << 30 // 1GB

func IsImage(fileHeader *multipart.FileHeader) (bool, error) {
	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Read the first 512 bytes of the file
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return false, err
	}

	// Detect MIME type
	mimeType := http.DetectContentType(buffer)
	return mimeType == "image/jpeg" || mimeType == "image/png" || mimeType == "image/gif", nil
}

func IsVideo(fileHeader *multipart.FileHeader) (bool, error) {
	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Read the first 512 bytes of the file
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return false, err
	}

	// Detect MIME type
	mimeType := http.DetectContentType(buffer)
	validTypes := []string{"video/mp4", "video/quicktime", "video/x-msvideo", "video/x-matroska"}

	for _, t := range validTypes {
		if mimeType == t {
			return true, nil
		}
	}
	return false, nil
}

func MakeUniqueIDWithTime() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func GetFileExtension(fileHeader *multipart.FileHeader) string {
	return filepath.Ext(fileHeader.Filename)
}

//func GetFileUrl(fileUrl string, rootPath string, folderPath string, fileName string) string {
//	return fmt.Sprintf("%s%s%s%s", fileUrl, "/api/file",
//		strings.Replace(folderPath, rootPath, "", 1), fileName)
//}
//
//func EnsureDir(dirPath string) error {
//	err := os.MkdirAll(dirPath, os.ModePerm)
//	if err != nil {
//		slog.Error("Error create dir", "dirPath", dirPath, "err", err.Error())
//		return err
//	}
//	return nil
//}
//
//func SaveImage(file *multipart.FileHeader, folderPath string) (string, error) {
//	if file.Size > MaxImageSize {
//		return "", errors.New("file size exceeds the maximum allowed limit of 1MB")
//	}
//
//	isImage, err := IsImage(file)
//	if err != nil {
//		return "", err
//	}
//
//	if !isImage {
//		return "", errors.New("file is not an image")
//	}
//
//	// save image
//	name := MakeUniqueIDWithTime()
//	path := fmt.Sprintf("%s/%s%s", folderPath, name, GetFileExtension(file))
//
//	src, err := file.Open()
//	if err != nil {
//		return "", err
//	}
//	defer src.Close()
//
//	dst, err := os.Create(path)
//	if err != nil {
//		return "", err
//	}
//	defer dst.Close()
//
//	if _, err = io.Copy(dst, src); err != nil {
//		if err := os.Remove(path); err != nil {
//			log.Println(err)
//		}
//		return "", err
//	}
//
//	return fmt.Sprintf("%s%s", name, GetFileExtension(file)), nil
//}
//
//func InitFolder(conf Config) error {
//	paths := reflect.ValueOf(&conf.FileStorage).Elem()
//
//	for i := 0; i < paths.NumField(); i++ {
//		fieldValue := paths.Field(i).String()
//
//		if fieldValue != "" {
//			if err := os.MkdirAll(fieldValue, 0755); err != nil {
//				return err
//			}
//		}
//	}
//
//	return nil
//}
