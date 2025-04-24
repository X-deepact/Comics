package handler

import (
	"be-client/internal/biz"
	"be-client/internal/dto/request"
	"be-client/util"
	"context"
	"errors"
	"pkg-common/common"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ComicHandler interface {
	GetPopularComicsByClassification(c *fiber.Ctx) error
	GetCommicDetailById(c *fiber.Ctx) error
	GetComicsBySearch(c *fiber.Ctx) error
	GetComicsByRecently(c *fiber.Ctx) error
	GetComicAndDramaBySearch(c *fiber.Ctx) error
}

type comicHandler struct {
	cu biz.ComicBiz
}

func NewComicHandler(cu biz.ComicBiz) ComicHandler {
	return &comicHandler{cu}
}

// GetPopularComicsByClassification fetches popular comics based on filters
// @Summary Get popular comics
// @Description Retrieve a list of popular comics filtered by genre, progress, audience, language, original_language, page, and page size.
// @Tags Comics
// @Accept json
// @Produce json
// @Param genre_id query int false "Filter by genre ID" example(1)
// @Param progress query string false "Filter by progress (ongoing, completed)" example("completed")
// @Param audience query string false "Filter by audience (teen, adult)" example("teen")
// @Param language query string false "Filter by language (en,zh,vi, etc.)" example("en")
// @Param original_language query string false "Filter by Original Language (en,zh,vi, etc.)" example("en")
// @Param page query int false "Page number for pagination" default(1) example(1)
// @Param page_size query int false "Number of items per page" default(10) example(10)
// @Success 200 {object} common.Response{data=[]response.ComicWithChapterResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/comic [get]
func (h *comicHandler) GetPopularComicsByClassification(c *fiber.Ctx) error {
	var req request.GetPopularComicsRequest

	if err := c.QueryParser(&req); err != nil {
		return util.ResponseApiPagination(c, nil, &common.Pagination{}, errors.New("Invalid query parameters"))
	}

	allQueries := c.Queries()

	allowedKeys := map[string]bool{
		"genre_id":          true,
		"progress":          true,
		"audience":          true,
		"language":          true,
		"original_language": true,
		"page_size":         true,
		"page":              true,
	}

	for key := range allQueries {
		if !allowedKeys[key] {
			return util.ResponseApiPagination(c, nil, &common.Pagination{}, errors.New("Invalid query parameters"))
		}
	}

	req.Genre, _ = strconv.Atoi(c.Query("genre_id", "0"))
	req.Page, _ = strconv.Atoi(c.Query("page", "1"))
	req.PageSize, _ = strconv.Atoi(c.Query("page_size", "10"))

	res, pagination, err := h.cu.GetPopularComicsByClassification(c.Context(), req)
	if err != nil {
		return util.ResponseApiPagination(c, nil, &pagination, err)
	}
	return util.ResponseApiPagination(c, res, &pagination, nil)
}

// GetCommicDetailByCode fetches comic detail by id
// @Summary Get comic detail by id
// @Description Retrieve a comic detail by id
// @Tags Comics
// @Accept json
// @Produce json
// @Param id path string true "Comic id" example("id")
// @Success 200 {object} common.Response{data=response.ComicDetailResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/comic/{id} [get]
func (h *comicHandler) GetCommicDetailById(c *fiber.Ctx) error {
	code := c.Params("id")
	if code == "" {
		return util.ResponseApi(c, nil, errors.New("Bad Request"))
	}

	languge := util.GetLanguage(c)
	id := util.StringToInt64(code)

	res, err := h.cu.GetComicDetailById(c.Context(), id, languge)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}
	return util.ResponseApi(c, res, nil)
}

// GetComicsBySearch Retrieve a list of comics, optionally filtered by language and search_keyword(comic name or author name)
// @Summary Get all Recommend Comics
// @Description Retrieve a list of comics by search
// @Tags Comics
// @Accept json
// @Produce json
// @Param language query string false "Filter by language (en,cn,vi, etc.)" example("en")
// @Param search_keyword query string false "Filter by comic name or author name (Superhero Adventures, etc.)" example("Superhero Adventures")
// @Param page query int false "Page number for pagination" default(1) example(1)
// @Param page_size query int false "Number of items per page" default(10) example(10)
// @Success 200 {object} common.Response{data=[]response.ComicWithChapterAndAuthorResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/comic/search [get]
func (h *comicHandler) GetComicsBySearch(c *fiber.Ctx) error {
	var req request.GetComicSearchRequest

	if err := c.QueryParser(&req); err != nil {
		return util.ResponseApiPagination(c, nil, &common.Pagination{}, errors.New("Invalid query parameters"))
	}

	allQueries := c.Queries()

	allowedKeys := map[string]bool{
		"search_keyword": true,
		"language":       true,
		"page_size":      true,
		"page":           true,
	}

	for key := range allQueries {
		if !allowedKeys[key] {
			return util.ResponseApiPagination(c, nil, &common.Pagination{}, errors.New("Invalid query parameters"))
		}
	}

	res, pagination, err := h.cu.GetComicsBySearch(context.Background(), req)
	if err != nil {
		return util.ResponseApiPagination(c, nil, &pagination, err)
	}
	return util.ResponseApiPagination(c, res, &pagination, nil)
}

// GetComicsByRecently Retrieve a list of recently comics(limit 10)
// @Summary Get recently comics
// @Description Retrieve a list of recently comics
// @Tags Comics
// @Accept json
// @Produce json
// @Param language query string false "Filter by language (en,zh,vi, etc.)" example("en")
// @Param page query int false "Page number for pagination" default(1) example(1)
// @Param page_size query int false "Number of items per page" default(10) example(10)
// @Success 200 {object} common.Response{data=[]response.ComicRecent} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/comic/recent [get]
func (h *comicHandler) GetComicsByRecently(c *fiber.Ctx) error {
	var req request.GetComicRecentRequest

	if err := c.QueryParser(&req); err != nil {
		return util.ResponseApiPagination(c, nil, &common.Pagination{}, errors.New("Invalid query parameters"))
	}

	allQueries := c.Queries()

	allowedKeys := map[string]bool{
		"page_size": true,
		"page":      true,
		"language":  true,
	}

	for key := range allQueries {
		if !allowedKeys[key] {
			return util.ResponseApiPagination(c, nil, &common.Pagination{}, errors.New("Invalid query parameters"))
		}
	}

	res, pagination, err := h.cu.GetComicsByRecently(context.Background(), req)
	if err != nil {
		return util.ResponseApiPagination(c, nil, &pagination, err)
	}
	return util.ResponseApiPagination(c, res, &pagination, nil)
}

// GetComicAndDramaBySearch Retrieve a list of comics and drama, optionally filtered by language and search_keyword(comic name or author name or drama name)
// @Summary Get all comics and drama
// @Description Retrieve a list of comics and drama by search
// @Tags Comics And Short Drama
// @Accept json
// @Produce json
// @Param language query string false "Filter by language (en,cn,vi, etc.)" example("en")
// @Param search_keyword query string false "Filter by comic name or author name or drama name (Superhero Adventures, etc.)" example("Superhero Adventures")
// @Param page query int false "Page number for pagination" default(1) example(1)
// @Param page_size query int false "Number of items per page" default(10) example(10)
// @Success 200 {object} common.Response{data=[]response.ComicAndDramaSearchResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/comic/searchAll [get]
func (h *comicHandler) GetComicAndDramaBySearch(c *fiber.Ctx) error {
	var req request.GetComicSearchRequest

	if err := c.QueryParser(&req); err != nil {
		return util.ResponseApiPagination(c, nil, &common.Pagination{}, errors.New("Invalid query parameters"))
	}

	allQueries := c.Queries()

	allowedKeys := map[string]bool{
		"search_keyword": true,
		"language":       true,
		"page_size":      true,
		"page":           true,
	}

	for key := range allQueries {
		if !allowedKeys[key] {
			return util.ResponseApiPagination(c, nil, &common.Pagination{}, errors.New("Invalid query parameters"))
		}
	}

	res, pagination, err := h.cu.GetComicAndDramaBySearch(context.Background(), req)
	if err != nil {
		return util.ResponseApiPagination(c, nil, &pagination, err)
	}
	return util.ResponseApiPagination(c, res, &pagination, nil)
}
