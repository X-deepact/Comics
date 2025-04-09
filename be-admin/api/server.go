package api

import (
	"comics-admin/db"
	"comics-admin/token"
	config "comics-admin/util"
	"comics-admin/val"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/minio/minio-go/v7"

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
	minio      *config.MinioConfig
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

	server.setUpMinio()

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
	s.genreDramaRouter()
	s.generalRouter()
	s.authorRoutes()
	s.recommendRoutes()
	s.chapterRouter()
	s.uploadDramaRouter()
	s.shortDramaRoutes()
	s.episodeRoutes()
	//router.POST("/register", s.register)
	//router.POST("/login", s.login)

	//authRoutes := router.Group("/").Use(s.authMiddleware(s.tokenMaker))

	// Ads routes
	adsGroup := s.router.Group("/api/ads")
	adsGroup.Use(s.authMiddleware(s.tokenMaker))
	{
		adsGroup.POST("", s.createAds)
		adsGroup.GET("", s.getAdsList)
		adsGroup.GET("/:id", s.getAdsById)
		adsGroup.PUT("", s.updateAds)
		adsGroup.DELETE("/:id", s.deleteAds)
		adsGroup.PATCH("/:id/status", s.updateAdsStatus)
		adsGroup.POST("/upload-image", s.uploadAdsImage)
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

func (r *Server) GetUserIdFromContext(ctx *gin.Context) (int64, error) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		return 0, fmt.Errorf("user not authenticated")
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		return 0, fmt.Errorf("invalid user ID type")
	}

	return userIDInt64, nil
}

func (r *Server) GetUsernameFromContext(ctx *gin.Context) (string, error) {
	issuer, exists := ctx.Get("username")
	if !exists {
		return "", fmt.Errorf("claims not found")
	}
	return issuer.(string), nil
}

func (s *Server) setUpMinio() {
	minioConfig := config.NewMinioClient(
		s.config.FileStorage.Endpoint,
		s.config.FileStorage.AccessKey,
		s.config.FileStorage.SecretKey,
		s.config.FileStorage.BucketName,
		s.config.FileStorage.UseSSL,
		s.config.ApiFile.Url,
	)

	// Set bucket public
	err := setBucketPublic(minioConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to set bucket policy")
	}

	s.minio = minioConfig
}

func setBucketPublic(minioConfig *config.MinioConfig) error {
	// Check bucket exists
	ctx := context.Background()
	exists, err := minioConfig.Client.BucketExists(ctx, minioConfig.BucketName)
	if err != nil {
		return err
	}

	if !exists {
		if err = minioConfig.Client.MakeBucket(ctx, minioConfig.BucketName, minio.MakeBucketOptions{}); err != nil {
			return err
		}
	}

	policy := fmt.Sprintf(`{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Principal": "*",
                "Action": "s3:GetObject",
                "Resource": "arn:aws:s3:::%s/*"
            }
        ]
    }`, minioConfig.BucketName)

	return minioConfig.Client.SetBucketPolicy(ctx, minioConfig.BucketName, policy)
}

func ExtractUserID(ctx *gin.Context) (int64, error) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		return 0, fmt.Errorf("user not authenticated")
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		return 0, fmt.Errorf("invalid user ID type")
	}

	return userIDInt64, nil
}
