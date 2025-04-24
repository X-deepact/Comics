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

func SetupEpisodeRoute(r fiber.Router, db database.DBInterface, config *config.Config) {
	episodeRepo := repository.NewEpisodeRepository(db)
	episodeMapper := mapper.NewEpisodeMapper(config)
	episodeUc := biz.NewEpisodeBiz(episodeRepo, episodeMapper)
	episodeHandler := handler.NewEpisodeHandler(episodeUc)

	groupComic := r.Group("/episode")
	GET(groupComic, "/:id", episodeHandler.GetEpisodeById)
}
