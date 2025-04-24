package config

import (
	"fmt"
	"log"
	"log/slog"
	"net/url"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server      ServerConfig      `mapstructure:"server"`
	FileStorage FileStorageConfig `mapstructure:"file_storage"`
	Profile     ProfileConfig     `mapstructure:"profile"`
	Database    DatabaseConfig    `mapstructure:"database"`
	Log         LogConfig         `mapstructure:"log"`
	Authen      AuthenConfig      `mapstructure:"auth"`
}

type ServerConfig struct {
	AppName      string `mapstructure:"APP_NAME"`
	Port         int    `mapstructure:"PORT"`
	Address      string `mapstructure:"ADDRESS"`
	ReadTimeout  int    `mapstructure:"READ_TIMEOUT"`
	WriteTimeout int    `mapstructure:"WRITE_TIMEOUT"`
}

type ProfileConfig struct {
	Env string `mapstructure:"ENV"`
}

type DatabaseConfig struct {
	DriverName         string        `mapstructure:"DRIVER_NAME"`
	Host               string        `mapstructure:"HOST"`
	Port               int           `mapstructure:"PORT"`
	UserName           string        `mapstructure:"USERNAME"`
	Password           string        `mapstructure:"PASSWORD"`
	DBName             string        `mapstructure:"DB_NAME"`
	SSLMode            string        `mapstructure:"SSL_MODE"`
	MaxOpenConnections int           `mapstructure:"MAX_OPEN_CONNECTIONS"`
	MaxIdleConnections int           `mapstructure:"MAX_IDLE_CONNECTIONS"`
	MaxConnLifetime    time.Duration `mapstructure:"MAX_CONN_LIFETIME"`
	MaxConnIdleTime    time.Duration `mapstructure:"MAX_CONN_IDLE_TIME"`
}

type LogConfig struct {
	LogLevel   slog.Level `mapstructure:"LEVEL"`
	JSONOutput bool       `mapstructure:"JSON_OUTPUT"`
	AddSource  bool       `mapstructure:"ADD_SOURCE"`
}

type AuthenConfig struct {
	PasswordSalt           string        `mapstructure:"PASSWORD_SALT"`
	AccessSecret           string        `mapstructure:"ACCESS_SECRET"`
	RefreshSecret          string        `mapstructure:"REFRESH_SECRET"`
	AccessTokenExpiration  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRATION"`
	RefreshTokenExpiration time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRATION"`
}

type FileStorageConfig struct {
	EndPoint            string `mapstructure:"ENDPOINT"`
	FileUrl             string `mapstructure:"FILE_URL"`
	AccessKey           string `mapstructure:"ACCESS_KEY"`
	SecretKey           string `mapstructure:"SECRET_KEY"`
	BucketName          string `mapstructure:"BUCKET_NAME"`
	UseSSL              bool   `mapstructure:"USE_SSL"`
	AvatarFilePath      string `mapstructure:"AVATAR_FILE_PATH"`
	ComicCoverFilePath  string `mapstructure:"COMIC_COVER_FILE_PATH"`
	ChapterItemFilePath string `mapstructure:"CHAPTER_ITEM_FILE_PATH"`
	AdsFilePath         string `mapstructure:"ADS_FILE_PATH"`
	RecommendFilePath   string `mapstructure:"RECOMMEND_FILE_PATH"`
	DramaThumbFilePath  string `mapstructure:"SHORT_DRAMA_THUMB_FILE_PATH"`
	DramaSubFilePath    string `mapstructure:"SHORT_DRAMA_SUB_FILE_PATH"`
	DramaFilePath       string `mapstructure:"SHORT_DRAMA_FILE_PATH"`
}

func loadConfig(configfile string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigFile(configfile)
	v.SetConfigType("ini")
	v.AutomaticEnv()

	v.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	if err := v.ReadInConfig(); err != nil {
		log.Println("config file not found. Using exists environment variables")
		return nil, err
	}
	return v, nil
}

func LoadConfig(pathToFile string, config any) error {
	configPath := pathToFile + "/" + "config"
	confFile := fmt.Sprintf("%s.ini", configPath)
	slog.Info(fmt.Sprintf("Config file path: %s", confFile))
	v, err := loadConfig(confFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := v.Unmarshal(&config); err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
	}
	return nil
}

func (r *DatabaseConfig) BuildDnsMYSQL() string {
	loc := url.QueryEscape("Asia/Shanghai")
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s", r.UserName, r.Password, r.Host, r.Port, r.DBName, loc)
}
