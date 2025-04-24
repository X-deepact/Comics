package route

import (
	"be-client/config"
	"be-client/internal/infra/database"
	"be-client/internal/infra/http"
	"be-client/internal/middleware"
)

func InitRoute(cf *config.Config, db database.DBInterface) {
	httpClient := http.HttpServer{
		AppName: cf.Server.AppName,
		Conf:    &cf.Server,
	}
	httpClient.InitHttpServer()
	apiv1 := httpClient.App().Group("/api/v1")

	auth := middleware.NewAuthenHandler(cf.Authen)

	SetupUserRoute(apiv1, auth, db, cf)
	SetupChapterRoute(apiv1, db, cf)

	SetupGenreRoute(apiv1, db, cf)
	SetupComicRoute(apiv1, db, cf)
	SetupRecommendComicRoute(apiv1, db, cf)
	SetupAdsRoute(apiv1, db, cf)
	SetupSubtitleRoute(apiv1, db, cf)
	SetupShortDramaRoute(apiv1, db, cf)
	SetupGenreForShortDramaRoute(apiv1, db, cf)
	SetupEpisodeRoute(apiv1, db, cf)
	httpClient.Start()
}
