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

func SetupRecommendComicRoute(r fiber.Router, db database.DBInterface, config *config.Config) {
	recommendComicRepo := repository.NewRecommendComicRepository(db)
	recommendRepo := repository.NewRecommendManagerRepository(db)
	chapterRepo := repository.NewChapterRepository(db)

	chapterMapper := mapper.NewChapterMapper(config)
	recommendComicMapper := mapper.NewRecommendComicMapper(config)

	recommendComicUc := biz.NewRecommendComicBiz(recommendComicRepo, recommendComicMapper, recommendRepo, chapterRepo, chapterMapper)
	recommendComicHandler := handler.NewRecommendComicHandler(recommendComicUc)

	groupRecommend := r.Group("/recommend-comics")
	GET(groupRecommend, "/all", recommendComicHandler.GetRecommendComics)
}
