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

func SetupComicRoute(r fiber.Router, db database.DBInterface, config *config.Config) {
	comicRepo := repository.NewComicRepository(db)
	chapterRepo := repository.NewChapterRepository(db)

	comicMapper := mapper.NewComicMapper(config)
	chapterMapper := mapper.NewChapterMapper(config)

	authorMapper := mapper.NewAuthorMapper(config)
	authorRepo := repository.NewAuthorRepository(db)
	comicAuthorRepo := repository.NewComicAuthorRepository(db)

	genreMapper := mapper.NewGenreMapper(config)
	genreRepo := repository.NewGenreRepository(db)

	shortDramaRepository := repository.NewShortDramaRepository(db)
	dramaTranslationRepo := repository.NewDramaTranslationRepository(db)
	episodeRepo := repository.NewEpisodeRepository(db)

	comicUc := biz.NewComicBiz(comicRepo, comicMapper,
		chapterRepo, chapterMapper,
		authorRepo, comicAuthorRepo, authorMapper,
		genreRepo, genreMapper, shortDramaRepository, dramaTranslationRepo, episodeRepo)
	comicHandler := handler.NewComicHandler(comicUc)

	groupComic := r.Group("/comic")
	GET(groupComic, "/search", comicHandler.GetComicsBySearch)
	GET(groupComic, "/recent", comicHandler.GetComicsByRecently)
	GET(groupComic, "", comicHandler.GetPopularComicsByClassification)
	GET(groupComic, "/searchAll", comicHandler.GetComicAndDramaBySearch)
	GET(groupComic, "/:id", comicHandler.GetCommicDetailById)

}
