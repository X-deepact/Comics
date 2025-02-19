package api

import (
	"comics-admin/db"
	"comics-admin/token"
	config "comics-admin/util"
	"comics-admin/val"
	"fmt"
	"net/http"
	"os"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(cors.New(cors.Config{
		AllowOrigins: s.config.Web.AllowedOrigins, // Allowed domains
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowHeaders: []string{"*"}, // Allow all headers
	}))

	fileH := router.Group("/api/file")
	fileH.Static("/", s.config.FileStorage.RootFolder)

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

	s.router = router

	s.userRouter()
	s.comicRouter()
	s.genreRouter()
	s.generalRouter()
	s.authorRoutes()
	s.recommendRoutes()
	s.chapterRouter()
	//router.POST("/register", s.register)
	//router.POST("/login", s.login)

	//authRoutes := router.Group("/").Use(s.authMiddleware(s.tokenMaker))

	/*
		authRoutes.GET("/comics", s.getAllSiteItems)
		authRoutes.GET("/comic/:id", s.getSiteItem)
		authRoutes.DELETE("/comic/:id", s.deleteSiteItem)
		authRoutes.DELETE("/comics", s.deleteAllSiteItems)
	*/
	//s.router = router

	// Ads routes
	adsGroup := s.router.Group("/api/ads")
	adsGroup.Use(s.authMiddleware(s.tokenMaker))
	{
		adsGroup.POST("", s.createAds)
		adsGroup.GET("", s.getAdsList)
		adsGroup.PUT("", s.updateAds)
		adsGroup.DELETE("/:id", s.deleteAds)
	}

	chapterItemGroup := s.router.Group("/api/chapter-items")
	chapterItemGroup.Use(s.authMiddleware(s.tokenMaker))
	{
		chapterItemGroup.POST("", s.createChapterItem)
		chapterItemGroup.GET("/:id", s.getChapterItem)
		chapterItemGroup.GET("", s.listChapterItems)
		chapterItemGroup.PUT("", s.updateChapterItem)
		chapterItemGroup.DELETE("/:id", s.deleteChapterItem)
		chapterItemGroup.POST("/upload-image", s.uploadChapterImage)
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func errorResponseMessage(message string) gin.H {
	return gin.H{"error": message}
}
