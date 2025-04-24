package route

import (
	"be-client/config"
	"be-client/internal/biz"
	"be-client/internal/handler"
	"be-client/internal/infra/database"
	"be-client/internal/mapper"
	"be-client/internal/middleware"
	"be-client/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoute(r fiber.Router, auth *middleware.AuthenHandler, db database.DBInterface, config *config.Config) {
	userRepo := repository.NewUserRepository(db)
	userMapper := mapper.NewUserMapper()

	profileRepo := repository.NewProfileRepository(db)
	profileMapper := mapper.NewProfileMapper(config)

	tierRepo := repository.NewTierRepository(db)

	userUc := biz.NewUserBiz(userRepo, profileRepo, tierRepo, profileMapper, userMapper, *auth)
	userHandler := handler.NewUserHandler(userUc, *auth)

	authGroup := r.Group("/auth")
	authHandler := handler.NewAuthHandler(userUc, *auth)
	POST(authGroup, "/register", authHandler.RegisterUser)
	POST(authGroup, "/login", authHandler.LoginUser)
	POST(authGroup, "/refresh", authHandler.RefreshToken)

	groupUser := r.Group("/user")
	protectedRoute := groupUser.Use(auth.AuthMiddleware())
	GET(protectedRoute, "/profile", userHandler.Profile)
	PUT(protectedRoute, "/update", userHandler.UpdateUser)
	PUT(protectedRoute, "/change-password", userHandler.ChangePassword)
}
