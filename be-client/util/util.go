package util

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"strconv"
	"time"

	"pkg-common/common"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

const (
	ContextTimeout     = 30 * time.Second
	XHeaderLanguage    = "X-Language"
	DefaultTierNewUser = 1
)

func GetLanguage(c *fiber.Ctx) string {
	lang := c.Get(XHeaderLanguage)
	if len(lang) == 0 {
		return "en"
	}
	return lang
}

func StringToInt64(s string) int64 {
	fmt.Println(s)
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found. Using exists environment variables")
	}
}

func UUID() string {
	id := uuid.New()
	return id.String()
}

func UUIDFunc() func() string {
	return func() string {
		return UUID()
	}
}

func ContextwithTimeout() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), ContextTimeout)
	go cancelContext(ContextTimeout, cancel)
	return ctx
}

func cancelContext(timeout time.Duration, cancel context.CancelFunc) {
	time.Sleep(timeout)
	cancel()
	slog.Info("context canceled by timeout")
}

func ToInt64(value interface{}) int64 {
	switch v := value.(type) {
	case int64:
		return v
	case int:
		return int64(v)
	case int32:
		return int64(v)
	case uint:
		return int64(v)
	case uint64:
		return int64(v)
	case float64:
		return int64(v)
	case json.Number:
		i, err := v.Int64()
		if err == nil {
			return i
		}
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			return i
		}
	default:
		return 0
	}
	return 0
}

func ToTimePtr(value interface{}) *time.Time {
	switch v := value.(type) {
	case time.Time:
		return &v
	case *time.Time:
		return v
	case string:
		parsedTime, err := time.Parse("2006-01-02 15:04:05", v)
		if err == nil {
			return &parsedTime
		}
	}
	return nil
}

func GetStringOrEmpty(val interface{}) string {
	if val == nil {
		return ""
	}

	switch v := val.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	default:
		return ""
	}
}

func GetFileUrl(fileUrl, bucketName, folderPath, fileName string) string {
	if fileName == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s%s%s", fileUrl, bucketName, folderPath, fileName)
}

func ResponseApi(c *fiber.Ctx, data any, err error) error {
	apiResponse := common.ApiResponse(data, nil, err)
	return c.Status(fiber.StatusOK).JSON(apiResponse)
}

func ResponseApiPagination(c *fiber.Ctx, data any, pagin *common.Pagination, err error) error {
	apiResponse := common.ApiResponse(data, pagin, err)
	return c.Status(fiber.StatusOK).JSON(apiResponse)
}

func GetDefaultLanguage() string {
	return "en"
}
