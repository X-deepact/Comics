package handler

import (
	"be-client/internal/biz"
	"be-client/util"

	"github.com/gofiber/fiber/v2"
)

type GenreHandler interface {
	GetAllGenres(c *fiber.Ctx) error
}

type genreHandler struct {
	gu biz.GenreBiz
}

func NewGenreHandler(gu biz.GenreBiz) GenreHandler {
	return &genreHandler{gu}
}

// GetAllGenres retrieves genres, optionally filtering for homepage and language
// @Summary Get all genres
// @Description Retrieve a list of all genres, or only for the homepage, with optional language filtering
// @Tags Genres
// @Accept json
// @Produce json
// @Param isHomePage query bool false "Set to true to get only homepage genres (position 1-15)"
// @Param language query string false "Language code for filtering genres (e.g., 'en', 'fr', 'es')"
// @Success 200 {object} common.Response{data=[]response.GenreResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/genre/all [get]
func (h *genreHandler) GetAllGenres(c *fiber.Ctx) error {
	isHomePage := c.Query("isHomePage") == "true"
	lang := c.Query("language")

	resp, err := h.gu.GetAllGenres(isHomePage, lang)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}

	return util.ResponseApi(c, resp, nil)
}
