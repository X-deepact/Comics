package main

import (
	"comics-admin/api"
	"comics-admin/db"
	_ "comics-admin/docs"
	config "comics-admin/util"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// @title          			   Admin Comics API Documentation
// @version         		   1.0
// @description     		   Swagger Admin Comics API Documentation.
// @host            		   localhost:8080
// @BasePath       			   /
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
func main() {
	conf, err := config.LoadConfig("./config")

	if err != nil {
		log.Fatal().Err(err).Msg("Load Config Failed")
	}

	dsn := conf.Source.DBSource

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database")
	}

	fmt.Println("Successfully connected to the database!")

	store := db.NewStore(database)

	go runHttpServer(store, conf)

	listenForShutdown()

}

/*
	func runSQLSchema(db *gorm.DB) error {
		schema, err := ioutil.ReadFile("schema.sql")
		if err != nil {
			return err
		}

		statements := strings.Split(string(schema), ";")

		for _, stmt := range statements {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}

			if err := db.Exec(stmt).Error; err != nil {
				return fmt.Errorf("failed to execute SQL statement: %v\nStatement: %s", err, stmt)
			}
		}

		return nil
	}
*/
func runHttpServer(store db.Store, config config.Config) {

	s, err := api.NewServer(store, config)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create server")
	}

	err = s.Start(config.Source.ServerAdd)

	if err != nil {
		log.Fatal().Err(err).Msg("Cannot Start Sever..")
	}
}

func listenForShutdown() {
	var wg *sync.WaitGroup

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	shutdown(wg)
	os.Exit(0)
}
func shutdown(wg *sync.WaitGroup) {

	log.Info().Msg("would run clean up tasks...")

	wg.Wait()

	log.Info().Msg("closing channels and shutting down applications...")

}
