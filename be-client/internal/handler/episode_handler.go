package handler

import (
	"be-client/internal/biz"
	"be-client/util"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type EpisodeHandler interface {
	GetEpisodeById(c *fiber.Ctx) error
}

type episodeHandler struct {
	cu biz.EpisodeBiz
}

func NewEpisodeHandler(cu biz.EpisodeBiz) EpisodeHandler {
	return &episodeHandler{cu}
}

// GetEpisodeById fetches episode by id
// @Summary Get episode by id
// @Description Retrieve a episode by id
// @Tags Episode
// @Accept json
// @Produce json
// @Param id path string true "Episode id" example("id")
// @Success 200 {object} common.Response{data=response.EpisodeGetResponse} "Successful response or Fail response. Code can be 'SUCCESS' or 'ERROR'."
// @Router /v1/episode/{id} [get]
func (h *episodeHandler) GetEpisodeById(c *fiber.Ctx) error {
	code := c.Params("id")
	if code == "" {
		return util.ResponseApi(c, nil, errors.New("Bad Request"))
	}
	id := util.StringToInt64(code)

	res, err := h.cu.GetEpisodeById(c.Context(), id)
	if err != nil {
		return util.ResponseApi(c, nil, err)
	}
	return util.ResponseApi(c, res, nil)
}
