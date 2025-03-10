package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	DBSource                string        `mapstructure:"DB_SOURCE"`
	Env                     string        `mapstructure:"ENVIRONMENT"`
	ServerAdd               string        `mapstructure:"HTTP_SERVER_ADDR"`
	TokenSymmetricKey       string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AuthorizationTypeBearer string        `mapstructure:"AUTHORIZATION_TYPE_BEARER"`
	AuthorizationPayloadKey string        `mapstructure:"AUTHORIZATION_PAYLOAD_KEY"`
	AuthorizationHeaderName string        `mapstructure:"AUTHORIZATION_HEADER_NAME"`
	AccessTokenDuration     time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

type ApiFileConfig struct {
	Url string `mapstructure:"URL"`
}

type FileStorageConfig struct {
	Endpoint          string `mapstructure:"ENDPOINT"`
	AccessKey         string `mapstructure:"ACCESS_KEY"`
	SecretKey         string `mapstructure:"SECRET_KEY"`
	BucketName        string `mapstructure:"BUCKET_NAME"`
	UseSSL            bool   `mapstructure:"USE_SSL"`
	AvatarFolder      string `mapstructure:"AVATAR_FOLDER"`
	CoverFolder       string `mapstructure:"COMIC_COVER_FOLDER"`
	ChapterItemFolder string `mapstructure:"CHAPTER_ITEM_FOLDER"`
	RecommendFolder   string `mapstructure:"RECOMMEND_FOLDER"`
}

type WebConfig struct {
	AllowedOrigins []string `mapstructure:"ALLOWED_ORIGINS"`
}

type Config struct {
	Source      DatabaseConfig    `mapstructure:"database"`
	ApiFile     ApiFileConfig     `mapstructure:"api_file"`
	FileStorage FileStorageConfig `mapstructure:"file_storage"`
	Web         WebConfig         `mapstructure:"web"`
}

// LoadConfig loads the configuration from the specified path

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("config") // Config file name (without extension)
	viper.SetConfigType("ini")    // Config file format
	viper.AddConfigPath(path)     // Path to look for the config file
	viper.AutomaticEnv()          // Automatically map environment variables if present

	err = viper.ReadInConfig() // Read the config file
	if err != nil {
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	err = viper.Unmarshal(&config) // Map the config to the struct
	if err != nil {
		return config, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return config, nil
}
