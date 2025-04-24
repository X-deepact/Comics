package handler

import (
	"be-client/internal/biz"
	"be-client/internal/dto/response"
	"be-client/util"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ChapterHandler interface {
	GetChapterDetailById(ctx *fiber.Ctx) error
}

type chapterHandler struct {
	chapterBiz biz.ChapterBiz
}

func NewChapterHandler(chapterBiz biz.ChapterBiz) ChapterHandler {
	return &chapterHandler{chapterBiz: chapterBiz}
}

// @Summary Get chapter detail by id
// @Description Get chapter detail by id
// @Tags Chapters
// @Accept json
// @Produce json
// @Param id path int true "Chapter id"
// @Success 200 {object} common.Response{data=response.ChapterDetailResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/chapter/{id} [get]
func (h *chapterHandler) GetChapterDetailById(ctx *fiber.Ctx) error {
	chapterId := ctx.Params("id")
	if chapterId == "" {
		return util.ResponseApi(ctx, nil, errors.New(response.ErrInvalidRequest.Error))
	}

	languge := util.GetLanguage(ctx)

	idInt64, err := strconv.ParseInt(chapterId, 10, 64)
	if err != nil {
		return util.ResponseApi(ctx, nil, errors.New(response.ErrInvalidRequest.Error))
	}

	chapterDetail, err := h.chapterBiz.GetChapterDetailByChapterId(ctx.Context(), idInt64, languge)
	if err != nil {
		return util.ResponseApi(ctx, nil, err)
	}
	return util.ResponseApi(ctx, chapterDetail, nil)
}
