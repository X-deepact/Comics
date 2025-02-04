package api

import (
	"comics-admin/db"
	"comics-admin/token"
	config "comics-admin/util"
	"comics-admin/val"
	"fmt"
	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type Server struct {
	router     *gin.Engine
	store      db.Store
	config     config.Config
	tokenMaker token.Maker
}

func NewServer(store db.Store, config config.Config) (*Server, error) {
	maker, err := token.NewJWTMaker(config.Source.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}
	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: maker,
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("role", val.ValidateRole)

	}

	server.setUpRouter()

	return server, nil

}
func (s *Server) Start(address string) error {

	err := s.router.Run(address)

	if err != nil {

		return err
	}
	return nil

}
func (s *Server) setUpRouter() {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to set trusted proxies")
	}
	if s.config.Source.Env == "development" {
		router.Use(ginzerolog.Logger("GIN"))
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	router.POST("/register", s.register)
	router.POST("/login", s.login)

	//authRoutes := router.Group("/").Use(s.authMiddleware(s.tokenMaker))

	/*
		authRoutes.GET("/comics", s.getAllSiteItems)
		authRoutes.GET("/comic/:id", s.getSiteItem)
		authRoutes.DELETE("/comic/:id", s.deleteSiteItem)
		authRoutes.DELETE("/comics", s.deleteAllSiteItems)
	*/
	s.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
