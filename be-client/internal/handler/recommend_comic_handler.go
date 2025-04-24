package handler

import (
	"be-client/internal/biz"
	"be-client/internal/dto/request"
	"be-client/util"
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type RecommendComicHandler interface {
	GetRecommendComics(c *fiber.Ctx) error
}

type recommendComicHandler struct {
	recommendComicUc biz.RecommendComicBiz
}

func NewRecommendComicHandler(rc biz.RecommendComicBiz) RecommendComicHandler {
	return &recommendComicHandler{recommendComicUc: rc}
}

// GetRecommendComics retrieves recommend manager with comics, optionally filtering for language
// @Summary Get all Recommend Comics
// @Description Retrieve a list of all recommend manager with comics
// @Tags Recommend Comics
// @Accept json
// @Produce json
// @Param language query string false "Filter by language (en,zh,vi, etc.)" example("en")
// @Success 200 {object} common.Response{data=[]response.RecommendComicResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/recommend-comics/all [get]
func (h *recommendComicHandler) GetRecommendComics(c *fiber.Ctx) error {
	var req request.GetRecommendComicRequest
	if err := c.QueryParser(&req); err != nil {
		return util.ResponseApi(c, nil, errors.New("Invalid query parameters"))
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

	resp, err := h.recommendComicUc.GetRecommendComics(context.Background(), req)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}

	return util.ResponseApi(c, resp, nil)
}
