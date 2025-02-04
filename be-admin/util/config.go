package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
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

type Config struct {
	Source DatabaseConfig `mapstructure:"database"`
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
