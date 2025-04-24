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

func SetupSubtitleRoute(r fiber.Router, db database.DBInterface, config *config.Config) {
	subTitleRepo := repository.NewSubtitleRepository(db)

	subTitleMapper := mapper.NewSubtitleMapper(config)
	subTitleUc := biz.NewSubtitleBiz(subTitleMapper, subTitleRepo)
	subTitleHandler := handler.NewSubTitleHandler(subTitleUc)

	groupSubtitle := r.Group("/subtitles")
	GET(groupSubtitle, "/:episode_id", subTitleHandler.GetVTTs)
}
