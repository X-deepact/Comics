package handler

import (
	"be-client/internal/biz"
	"be-client/internal/dto/request"
	"be-client/util"
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type AdsHandler interface {
	GetAdsComics(c *fiber.Ctx) error
}

type adsHandler struct {
	cu biz.AdsBiz
}

func NewAdsHandler(cu biz.AdsBiz) AdsHandler {
	return &adsHandler{cu}
}

// GetAdsComics Retrieve a list of ads comics(limit 5 last updated date)
// @Summary Get ads comics
// @Description Retrieve a list of ads comics
// @Tags ads
// @Accept json
// @Produce json
// @Param language query string false "Filter by language (en,zh,vi, etc.)" example("en")
// @Success 200 {object} common.Response{data=[]response.AdsComicResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/ads/comic [get]
func (h *adsHandler) GetAdsComics(c *fiber.Ctx) error {
	var req request.GetAdsComicRequest

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

	resp, err := h.cu.GetAdsComics(context.Background(), req)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}

	return util.ResponseApi(c, resp, nil)
}
