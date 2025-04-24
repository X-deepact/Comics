package handler

import (
	"be-client/internal/biz"
	"be-client/internal/dto/request"
	"be-client/util"
	"context"
	"errors"
	"pkg-common/common"

	"github.com/gofiber/fiber/v2"
)

type GenreForShortDramaHandler interface {
	GetAllGenreForShortDrama(c *fiber.Ctx) error
}

type genreForShortDramaHandler struct {
	gu biz.GenreForShortDramaBiz
}

func NewGenreForShortDramaHandler(gu biz.GenreForShortDramaBiz) GenreForShortDramaHandler {
	return &genreForShortDramaHandler{gu}
}

// GetAllGenreForShortDrama retrieves genres for short drama, filtering for language
// @Summary Get all genres for short drama
// @Description Retrieve a list of all genres for short drama, with language filtering
// @Tags Genres For Short Drama
// @Accept json
// @Produce json
// @Param language query string false "Language code for filtering genres (e.g., 'en', 'fr', 'es')"
// @Success 200 {object} common.Response{data=[]response.GenresForShortDramaResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/short-drama-genre/all [get]
func (h *genreForShortDramaHandler) GetAllGenreForShortDrama(c *fiber.Ctx) error {
	var req request.GetGenresForShortDramaRequest

	if err := c.QueryParser(&req); err != nil {
		return util.ResponseApiPagination(c, nil, &common.Pagination{}, errors.New("Invalid query parameters"))
	}

	allQueries := c.Queries()

	allowedKeys := map[string]bool{
		"language": true,
	}

	for key := range allQueries {
		if !allowedKeys[key] {
			return util.ResponseApi(c, nil, errors.New("Invalid query parameters"))
		}
	}

	res, err := h.gu.GetAllGenreForShortDrama(context.Background(), req)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}
	return util.ResponseApi(c, res, nil)
}
