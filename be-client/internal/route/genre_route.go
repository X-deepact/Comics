package route

import (
	"be-client/config"
	"be-client/internal/biz"
	"be-client/internal/handler"
	"be-client/internal/infra/database"
	"be-client/internal/mapper"
	"be-client/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func SetupGenreRoute(r fiber.Router, db database.DBInterface, config *config.Config) {
	genreRepo := repository.NewGenreRepository(db)
	genreMapper := mapper.NewGenreMapper(config)
	genreUc := biz.NewGenreBiz(genreRepo, genreMapper)
	genreHandler := handler.NewGenreHandler(genreUc)

	groupGenre := r.Group("/genre")
	GET(groupGenre, "/all", genreHandler.GetAllGenres)
}
