package handler

import (
	"be-client/internal/biz"
	"be-client/internal/dto/request"
	"be-client/util"
	"errors"
	"pkg-common/common"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ShortDramaHandler interface {
	GetShortDramaByClassification(c *fiber.Ctx) error
	GetShortDramaDetailById(c *fiber.Ctx) error
}

type shortDramaHandler struct {
	cu biz.ShortDramaBiz
}

func NewShortDramaHandler(cu biz.ShortDramaBiz) ShortDramaHandler {
	return &shortDramaHandler{
		cu,
	}
}

// GetShortDramaByClassification fetches short drama based on filters
// @Summary Get short drama
// @Description Retrieve a list of short drama filtered by genre, release_year, language, page, and page size.
// @Tags Short Drama
// @Accept json
// @Produce json
// @Param genre_id query int false "Filter by genre ID" example(1)
// @Param release_year query string false "Filter by release year (2024, 2025)" example("2025")
// @Param language query string false "Filter by language (en,zh,vi, etc.)" example("en")
// @Param page query int false "Page number for pagination" default(1) example(1)
// @Param page_size query int false "Number of items per page" default(10) example(10)
// @Success 200 {object} common.Response{data=[]response.ShortDramaByClassificationResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/short-drama [get]
func (h *shortDramaHandler) GetShortDramaByClassification(c *fiber.Ctx) error {
	var req request.GetShortDramaByClassificationRequest

	if err := c.QueryParser(&req); err != nil {
		return util.ResponseApiPagination(c, nil, &common.Pagination{}, errors.New("Invalid query parameters"))
	}

	allQueries := c.Queries()

	allowedKeys := map[string]bool{
		"genre_id":     true,
		"release_year": true,
		"language":     true,
		"page_size":    true,
		"page":         true,
	}

	for key := range allQueries {
		if !allowedKeys[key] {
			return util.ResponseApiPagination(c, nil, &common.Pagination{}, errors.New("Invalid query parameters"))
		}
	}

	req.Genre, _ = strconv.Atoi(c.Query("genre_id", "0"))
	req.Page, _ = strconv.Atoi(c.Query("page", "1"))
	req.PageSize, _ = strconv.Atoi(c.Query("page_size", "10"))

	res, pagination, err := h.cu.GetShortDramaByClassification(c.Context(), req)
	if err != nil {
		return util.ResponseApiPagination(c, nil, &pagination, err)
	}
	return util.ResponseApiPagination(c, res, &pagination, nil)
}

// GetShortDramaDetailById fetches short drama detail by id
// @Summary Get short drama detail by id
// @Description Retrieve a short drama detail by id
// @Tags Short Drama
// @Accept json
// @Produce json
// @Param id path string true "Short Drama id" example("id")
// @Param language query string false "Filter by language (en,zh,vi, etc.)" example("en")
// @Success 200 {object} common.Response{data=response.ShortDramaDetailResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/short-drama/{id} [get]
func (h *shortDramaHandler) GetShortDramaDetailById(c *fiber.Ctx) error {
	var req request.GetShortDramaDetailRequest
	code := c.Params("id")
	if code == "" {
		return util.ResponseApi(c, nil, errors.New("Bad Request"))
	}
	id := util.StringToInt64(code)

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

	res, err := h.cu.GetShortDramaDetailById(c.Context(), id, req)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}
	return util.ResponseApi(c, res, nil)
}
