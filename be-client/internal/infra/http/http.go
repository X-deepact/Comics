package http

import (
	"be-client/config"
	"be-client/internal/middleware"
	"be-client/util"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type HttpServer struct {
	AppName string
	Conf    *config.ServerConfig
	app     *fiber.App
}

func (r *HttpServer) Start() {
	if r == nil || r.app == nil {
		slog.Error("http server is nil")
		return
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	serverShutdown := make(chan struct{})

	go func() {
		err := r.app.Listen(fmt.Sprintf("%s:%d", r.Conf.Address, r.Conf.Port))
		if err != nil {
			slog.Error("listen error", "err", err)
		}
		close(serverShutdown)
	}()

	<-shutdown
	slog.Info("Shutting down server...")

	ctx := util.ContextwithTimeout()
	if err := r.app.ShutdownWithContext(ctx); err != nil {
		slog.Info(fmt.Sprintf("Server forced to shutdown: %v\n", err))
	}

	<-serverShutdown
	slog.Info("Server gracefully stopped")

}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func (r *HttpServer) InitHttpServer() {
	app := fiber.New(r.ConfigFiber(r.Conf))
	app.Use(middleware.CorsFilter())
	SetupSwagger(app)
	r.app = app
}

func SetupSwagger(app *fiber.App) {
	app.Get("/swagger/*", swagger.New(swagger.Config{
		Title:        "Comics API Documentation",
		DeepLinking:  false,
		DocExpansion: "none",
	}))
}

func (r *HttpServer) App() *fiber.App {
	return r.app
}

func (r *HttpServer) ConfigFiber(conf *config.ServerConfig) fiber.Config {
	return fiber.Config{
		AppName:           conf.AppName,
		EnablePrintRoutes: true,
		ReadTimeout:       time.Duration(conf.ReadTimeout) * time.Second,
		WriteTimeout:      time.Duration(conf.WriteTimeout) * time.Second,
	}
}
