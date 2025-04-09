package api

import (
	"comics-admin/dto"
	config "comics-admin/util"
	"context"
	"errors"
	"pkg-common/common"
	"pkg-common/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) episodeRoutes() {
	group := s.router.Group("/api/episodes").Use(s.authMiddleware(s.tokenMaker))
	{
		group.POST("", s.createEpisode)
		group.PUT("", s.updateEpisode)
		group.GET("/:id", s.getEpisodeById)
		group.GET("", s.getEpisodes)
		group.DELETE("/:id", s.deleteEpisodeById)
	}
}

// @Summary Get list of episodes
// @Description Get list of episodes
// @Tags episodes
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param page_size query int true "Page size (must be between 10 and 50)" minimum(10) maximum(50)
// @Param sort_by query string false "Sort by"
// @Param sort query string false "Sort order"
// @Security BearerAuth
// @Success 200 {object} dto.ListSuccessResponse{data=[]dto.EpisodeResponse} "List episodes"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/episodes [get]
func (s *Server) getEpisodes(ctx *gin.Context) {
	req := &dto.EpisodeListRequest{}
	if err := req.Bind(ctx); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}
	episodes, total, err := s.store.ListEpisodes(req)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}
	userIds := []int64{}
	episodeIds := []int64{}
	for _, episode := range episodes {
		userIds = append(userIds, episode.CreatedBy, episode.UpdatedBy)
		episodeIds = append(episodeIds, episode.Id)
	}
	mapEpisodeIdSubtitle, err := s.store.GetSubtitleByEpisodeIds(context.Background(), episodeIds)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}
	mapIdUserName, err := s.store.GetUserNamesByIds(userIds)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	isSubtitleValid := len(mapEpisodeIdSubtitle) > 0
	resp := make([]*dto.EpisodeResponse, len(episodes))
	for i, episode := range episodes {
		resp[i] = s.mapFromEntityToResponse(episode)
		if episode.CreatedBy > 0 {
			resp[i].CreatedBy = mapIdUserName[episode.CreatedBy]
		}
		if episode.UpdatedBy > 0 {
			resp[i].UpdatedBy = mapIdUserName[episode.UpdatedBy]
		}
		if isSubtitleValid {
			resp[i].Subtitles = s.mapUrlSubtitleEntityToResponse(mapEpisodeIdSubtitle[episode.Id])
		}
	}
	config.BuildListResponse(ctx, &common.Pagination{Page: req.Page, PageSize: req.PageSize, Total: int(total)}, resp)
}

// @Summary Get an episode by ID
// @Description Get an episode by getEpisodeById
// @Tags episodes
// @Accept json
// @Produce json
// @Param id path int true "Episode ID"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.EpisodeResponse} "Episode retrieved successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/episodes/{id} [get]
func (s *Server) getEpisodeById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		config.BuildErrorResponse(ctx, errors.New("invalid id"), nil)
		return
	}

	episode, err := s.store.GetEpisodeById(id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}
	if episode == nil {
		config.BuildErrorResponse(ctx, errors.New("episode not found"), nil)
		return
	}
	resp := s.mapFromEntityToResponse(episode)
	mapIdUserName, err := s.store.GetUserNamesByIds([]int64{episode.CreatedBy, episode.UpdatedBy})
	if err == nil && len(mapIdUserName) > 0 {
		resp.CreatedBy = mapIdUserName[episode.CreatedBy]
		resp.UpdatedBy = mapIdUserName[episode.UpdatedBy]
	}
	subtitle, err := s.store.GetSubtitleByEpisodeId(context.Background(), resp.Id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}
	resp.Subtitles = s.mapUrlSubtitleEntityToResponse(subtitle)
	config.BuildSuccessResponse(ctx, resp)
}

// @Summary Create an createEpisode
// @Description Create an createEpisode
// @Tags episodes
// @Accept json
// @Produce json
// @Param episode body dto.EpisodeCreateRequest true "Episode Request"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.EpisodeResponse} "Episode created successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/episodes [post]
func (s *Server) createEpisode(ctx *gin.Context) {
	req := &dto.EpisodeCreateRequest{}
	if err := req.Bind(ctx); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	userId, err := s.GetUserIdFromContext(ctx)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}
	now := time.Now()
	episode := &model.EpisodeModel{
		DramaId:   req.DramaId,
		Number:    req.Number,
		Video:     req.Video,
		CreatedAt: &now,
		CreatedBy: userId,
		UpdatedAt: &now,
		Active:    config.Bool(false),
		UpdatedBy: userId,
	}

	err = s.store.CreateEpisode(episode)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	resp := s.mapFromEntityToResponse(episode)
	if len(req.Subtitles) > 0 {
		subtitles := make([]*model.SubtitleModel, len(req.Subtitles))
		for i, subtitle := range req.Subtitles {
			subtitles[i] = &model.SubtitleModel{
				EpisodeId:   episode.Id,
				Language:    subtitle.Language,
				SubtitleUrl: subtitle.Url,
			}
		}
		err := s.store.CreateSubtitles(context.Background(), subtitles)
		if err != nil {
			config.BuildErrorResponse(ctx, err, nil)
			return
		}
		resp.Subtitles = s.mapUrlSubtitleEntityToResponse(req.Subtitles)
	}

	mapIdUserName, err := s.store.GetUserNamesByIds([]int64{episode.CreatedBy, episode.UpdatedBy})
	if err == nil && len(mapIdUserName) > 0 {
		resp.CreatedBy = mapIdUserName[episode.CreatedBy]
		resp.UpdatedBy = mapIdUserName[episode.UpdatedBy]
	}
	config.BuildSuccessResponse(ctx, resp)
}

// @Summary Update an updateEpisode
// @Description Update an updateEpisode
// @Tags episodes
// @Accept json
// @Produce json
// @Param episode body dto.EpisodeUpdateRequest true "Episode Request"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse{data=dto.EpisodeResponse} "Episode updated successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/episodes [put]
func (s *Server) updateEpisode(ctx *gin.Context) {
	req := &dto.EpisodeUpdateRequest{}
	if err := req.Bind(ctx); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}
	if req.Id <= 0 {
		config.BuildErrorResponse(ctx, errors.New("invalid id"), nil)
		return
	}
	episode, err := s.store.GetEpisodeById(req.Id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}
	isUpdate := false
	if req.DramaId > 0 {
		isUpdate = true
		episode.DramaId = req.DramaId
	}
	if req.Number > 0 {
		isUpdate = true
		episode.Number = req.Number
	}
	if len(req.Video) > 0 {
		isUpdate = true
		episode.Video = req.Video
	}
	if req.Active != nil {
		isUpdate = true
		episode.Active = req.Active
	}
	if len(req.Subtitles) > 0 {
		isUpdate = true
	}
	now := time.Now()
	var resp *dto.EpisodeResponse
	if isUpdate {
		userId, err := s.GetUserIdFromContext(ctx)
		if err != nil {
			config.BuildErrorResponse(ctx, err, nil)
			return
		}
		episode.UpdatedAt = &now
		episode.UpdatedBy = userId
		err = s.store.UpdateEpisode(episode)
		if err != nil {
			config.BuildErrorResponse(ctx, err, nil)
			return
		}
		if len(req.Subtitles) > 0 {
			err = s.store.UpdateSubtitleUrlByEpisodeId(context.Background(), episode.Id, userId, req.Subtitles)
			if err != nil {
				config.BuildErrorResponse(ctx, err, nil)
				return
			}
		}
		resp = s.mapFromEntityToResponse(episode)
		mapIdUserName, err := s.store.GetUserNamesByIds([]int64{episode.CreatedBy, episode.UpdatedBy})
		if err == nil && len(mapIdUserName) > 0 {
			resp.CreatedBy = mapIdUserName[episode.CreatedBy]
			resp.UpdatedBy = mapIdUserName[episode.UpdatedBy]
		}
		subtitleResponse, err := s.store.GetSubtitleByEpisodeId(context.Background(), episode.Id)
		if err == nil && len(subtitleResponse) > 0 {
			resp.Subtitles = s.mapUrlSubtitleEntityToResponse(subtitleResponse)
		}
	}
	if resp != nil {
		config.BuildSuccessResponse(ctx, resp)
		return
	}
	config.BuildErrorResponse(ctx, errors.New("error update episode"), nil)
}

// @Summary Delete an episode by Id
// @Description Delete an episode by Id
// @Tags episodes
// @Accept json
// @Param id path int true "Episode ID"
// @Security BearerAuth
// @Success 200 {object} dto.SuccessResponse "Episode deleted successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/episodes/{id} [delete]
func (s *Server) deleteEpisodeById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		config.BuildErrorResponse(ctx, errors.New("invalid id"), nil)
		return
	}
	episode, err := s.store.GetEpisodeById(id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}
	if episode == nil {
		config.BuildErrorResponse(ctx, errors.New("episode not found"), nil)
		return
	}
	err = s.store.DeleteEpisode(id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	err = s.store.DeleteSubtitleByEpisodeId(context.Background(), id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}
	config.BuildSuccessResponse(ctx, nil)
}

func (s *Server) mapFromEntityToResponse(episode *model.EpisodeModel) *dto.EpisodeResponse {
	videoUrl := s.minio.GetFileUrl(s.config.FileStorage.ShortDramaFolder, episode.Video)
	resp := &dto.EpisodeResponse{
		Id:        episode.Id,
		DramaId:   episode.DramaId,
		Number:    episode.Number,
		Video:     videoUrl,
		Active:    episode.Active,
		CreatedAt: episode.CreatedAt.Format(time.RFC3339),
		UpdatedAt: episode.UpdatedAt.Format(time.RFC3339),
	}
	return resp
}

func (s *Server) mapUrlSubtitleEntityToResponse(subtitle []dto.Subtitle) []dto.Subtitle {
	if len(subtitle) == 0 {
		return []dto.Subtitle{}
	}
	resp := make([]dto.Subtitle, len(subtitle))
	for i, subtitle := range subtitle {
		resp[i] = dto.Subtitle{
			Language: subtitle.Language,
			Url:      s.minio.GetFileUrl(s.config.FileStorage.ShortDramaSubFolder, subtitle.Url),
		}
	}
	return resp
}