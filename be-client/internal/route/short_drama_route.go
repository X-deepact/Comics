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

func SetupShortDramaRoute(r fiber.Router, db database.DBInterface, config *config.Config) {
	shortDramaRepo := repository.NewShortDramaRepository(db)
	dramaTranslationRepo := repository.NewDramaTranslationRepository(db)
	episodeRepo := repository.NewEpisodeRepository(db)
	genreShortDramaRepo := repository.NewGenreForShortDramaRepository(db)

	shortDramaMapper := mapper.NewShortDramaMapper(config)
	shortDramaUc := biz.NewShortDramaBiz(shortDramaRepo, dramaTranslationRepo, shortDramaMapper, episodeRepo, genreShortDramaRepo)
	shortDramaHandler := handler.NewShortDramaHandler(shortDramaUc)

	groupShortDrama := r.Group("/short-drama")
	GET(groupShortDrama, "", shortDramaHandler.GetShortDramaByClassification)
	GET(groupShortDrama, "/:id", shortDramaHandler.GetShortDramaDetailById)
}
