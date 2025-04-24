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

func SetupAdsRoute(r fiber.Router, db database.DBInterface, config *config.Config) {
	adsRepo := repository.NewAdsRepository(db)
	adsMapper := mapper.NewAdsMapper(config)
	adsUc := biz.NewAdsBiz(adsRepo, adsMapper)
	adsHandler := handler.NewAdsHandler(adsUc)

	groupComic := r.Group("/ads")
	GET(groupComic, "/comic", adsHandler.GetAdsComics)
}
