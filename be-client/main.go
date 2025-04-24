package main

import (
	"be-client/config"
	"be-client/internal/infra/database"
	"be-client/internal/infra/logging"
	"be-client/internal/route"
	"be-client/util"
	"flag"
	"fmt"
	"log/slog"

	_ "be-client/docs"
)

var (
	flagConf   string
	configPath string
	envFile    string
)

func init() {
	util.LoadEnv()
	// Define flags
	flag.StringVar(&configPath, "config", "./config", "path to config file")
	flag.StringVar(&envFile, "env", "./config", "path to env file")
	flag.Parse()
}

func LoadConfig() *config.Config {
	conf := &config.Config{}
	err := config.LoadConfig(configPath, conf)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to load config, %v", err))
	}
	return conf
}

// @title Comics API
// @version 1.0
// @description Comics API Documentation
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@your-domain.com
// @host localhost:8080
// @BasePath /api
// @schemes http https

func main() {
	conf := LoadConfig()
	logging.InitLogger(*conf)
	StartServer(conf)
}

func StartServer(cfg *config.Config) {
	db := InitDatabase(cfg)
	route.InitRoute(cfg, db)
}

func InitDatabase(cfg *config.Config) database.DBInterface {
	db, err := database.NewDatabaseConnectClient(&cfg.Database)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to connect to database, %v", err))
		panic(err)
	}
	if db == nil {
		slog.Error("failed to connect to database")
		panic("failed to connect to database")
	}
	return db
}
