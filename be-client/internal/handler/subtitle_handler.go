package handler

import (
	"be-client/internal/biz"
	"be-client/util"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type SubtitleHandler interface {
	GetVTTs(c *fiber.Ctx) error
}

type subTitleHandler struct {
	cu biz.SubtitleBiz
}

func NewSubTitleHandler(cu biz.SubtitleBiz) SubtitleHandler {
	return &subTitleHandler{
		cu,
	}
}

// GetVTTs fetches VTTs by episode id
// @Summary Get VTTs by episode id
// @Description Retrieve a list of VTT by episode id
// @Tags Subtitles
// @Accept json
// @Produce json
// @Param episode_id path string true "Episode Id" example("episode_id")
// @Success 200 {object} common.Response{data=response.SubtitleResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/subtitles/{episode_id} [get]
func (h *subTitleHandler) GetVTTs(c *fiber.Ctx) error {
	episodeId, err := c.ParamsInt("episode_id")
	if err != nil {
		return util.ResponseApi(c, nil, errors.New("Invalid query parameters"))
	}

	vttUrls, err := h.cu.GetDramaVttFiles(c.Context(), int64(episodeId))
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}
	return util.ResponseApi(c, vttUrls, nil)
}
