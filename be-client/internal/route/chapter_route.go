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

func SetupChapterRoute(r fiber.Router, db database.DBInterface, config *config.Config) {
	comicRepo := repository.NewComicRepository(db)

	chapterRepo := repository.NewChapterRepository(db)
	chapterItemRepo := repository.NewChapterItemRepository(db)
	chapterMapper := mapper.NewChapterMapper(config)
	chapterBiz := biz.NewChapterBiz(chapterRepo, chapterItemRepo, chapterMapper, comicRepo)
	chapterHandler := handler.NewChapterHandler(chapterBiz)

	groupChapter := r.Group("/chapter")
	GET(groupChapter, "/:id", chapterHandler.GetChapterDetailById)
}
