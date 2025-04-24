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

func SetupGenreForShortDramaRoute(r fiber.Router, db database.DBInterface, config *config.Config) {
	genreForShortDramaRepo := repository.NewGenreForShortDramaRepository(db)
	genreForShortDramaMapper := mapper.NewGenreForShortDramaMapper(config)
	genreForShortDramaUc := biz.NewGenreForShortDramaBiz(genreForShortDramaRepo, genreForShortDramaMapper)
	genreForShortDramaHandler := handler.NewGenreForShortDramaHandler(genreForShortDramaUc)

	groupGenre := r.Group("/short-drama-genre")
	GET(groupGenre, "/all", genreForShortDramaHandler.GetAllGenreForShortDrama)
}
