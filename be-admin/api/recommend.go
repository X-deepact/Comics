package api

import (
	"comics-admin/dto"
	config "comics-admin/util"
	"errors"
	"pkg-common/common"
	"pkg-common/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) recommendRoutes() {
	group := s.router.Group("api/recommend").Use(s.authMiddleware(s.tokenMaker))
	group.POST("", s.CreateRecommend)
	group.GET("/:id", s.GetRecommendById)
	group.GET("", s.GetRecommends)
	group.PUT("", s.UpdateRecommendById)
	group.DELETE("/:id", s.DeleteRecommendById)
	group.POST("/comic", s.CreateRecommendComic)
	group.DELETE("/comic", s.DeleteRecommendComic)
}

// @Summary Create recommend
// @Description Create a new recommend
// @Tags recommends
// @Accept json
// @Produce json
// @Param     Authorization header string true "Bearer authorization token"
// @Param title formData string true "Title"
// @Param position formData int false "Position"
// @Param active_from formData int false "Active From"
// @Param active_to formData int false "Active To"
// @Param cover formData file false "Cover "
// @Security     BearerAuth
// @Success 200 {object} dto.ResponseMessage{data=dto.RecommendResponse} "Recommend created successfully"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/recommend [post]
func (s *Server) CreateRecommend(ctx *gin.Context) {
	var req dto.RecommendCreateRequest
	if err := req.Bind(ctx); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	file, _ := ctx.FormFile("cover")
	if file != nil {
		fileName, err := s.minio.SaveImage(file, s.config.FileStorage.RecommendFolder)
		if err != nil {
			config.BuildErrorResponse(ctx, err, nil)
			return
		}

		req.Cover = fileName
	}

	userId, err := s.GetUserIdFromContext(ctx)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	now := time.Now()
	activeFrom := time.UnixMilli(req.ActiveFrom)
	activeTo := time.UnixMilli(req.ActiveTo)
	recommendModel := model.RecommendManagerModel{
		Title:      req.Title,
		Cover:      req.Cover,
		Position:   req.Position,
		ActiveFrom: &activeFrom,
		ActiveTo:   &activeTo,
		CreatedAt:  &now,
		UpdatedAt:  &now,
		CreatedBy:  userId,
		UpdatedBy:  userId,
	}

	err = s.store.CreateRecommend(&recommendModel)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	resp := dto.RecommendResponse{
		Id:         recommendModel.Id,
		Title:      recommendModel.Title,
		Cover:      recommendModel.Cover,
		Position:   recommendModel.Position,
		ActiveFrom: recommendModel.ActiveFrom.Format(time.RFC3339),
		ActiveTo:   recommendModel.ActiveTo.Format(time.RFC3339),
		CreatedAt:  recommendModel.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  recommendModel.UpdatedAt.Format(time.RFC3339),
	}
	mapUserIdName, err := s.store.GetUserNamesByIds([]int64{recommendModel.CreatedBy, recommendModel.UpdatedBy})
	if err == nil && len(mapUserIdName) > 0 {
		resp.CreatedByName = mapUserIdName[recommendModel.CreatedBy]
		resp.UpdatedByName = mapUserIdName[recommendModel.UpdatedBy]
	}

	if resp.Cover != "" {
		resp.Cover = s.minio.GetFileUrl(s.config.FileStorage.RecommendFolder, resp.Cover)
	}

	config.BuildSuccessResponse(ctx, resp)
}

// @Summary Get recommend by Id
// @Description Get a recommend by Id
// @Tags recommends
// @Accept json
// @Produce json
// @Param     Authorization header string true "Bearer authorization token"
// @Param id path int true "Recommend ID"
// @Security     BearerAuth
// @Success 200 {object} dto.ResponseMessage{data=dto.RecommendResponse} "Recommend found"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/recommend/{id} [get]
func (s *Server) GetRecommendById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		config.BuildErrorResponse(ctx, errors.New("invalid id"), nil)
		return
	}

	recommend, err := s.store.GetRecommendById(id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	resp := dto.RecommendResponse{
		Id:         recommend.Id,
		Title:      recommend.Title,
		Position:   recommend.Position,
		ActiveFrom: recommend.ActiveFrom.Format(time.RFC3339),
		CreatedAt:  recommend.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  recommend.UpdatedAt.Format(time.RFC3339),
	}
	mapUserIdName, err := s.store.GetUserNamesByIds([]int64{recommend.CreatedBy, recommend.UpdatedBy})
	if err == nil && len(mapUserIdName) > 0 {
		resp.CreatedByName = mapUserIdName[recommend.CreatedBy]
		resp.UpdatedByName = mapUserIdName[recommend.UpdatedBy]
	}
	if recommend.ActiveTo != nil {
		resp.ActiveTo = recommend.ActiveTo.Format(time.RFC3339)
	}

	if recommend.Cover != "" {
		resp.Cover = s.minio.GetFileUrl(s.config.FileStorage.RecommendFolder, recommend.Cover)
	}

	config.BuildSuccessResponse(ctx, resp)
}

// @Summary List recommends
// @Description List all recommends
// @Tags recommends
// @Accept json
// @Produce json
// @Param     Authorization header string true "Bearer authorization token"
// @Param sort_by query string false "Sort By"
// @Param sort query string false "Sort"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param title query string false "Title"
// @Security     BearerAuth
// @Success 200 {object} dto.ResponseMessage{data=[]dto.RecommendResponse} "List recommends"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/recommend [get]
func (s *Server) GetRecommends(ctx *gin.Context) {
	var req dto.RequestQueryFilter
	if err := ctx.Bind(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	req.SortBy = config.GetSortBy(req.SortBy)
	req.Sort = config.GetSortOrder(req.Sort)

	recommends, total, err := s.store.GetRecommends(req)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}
	userIds := []int64{}
	for _, r := range recommends {
		userIds = append(userIds, r.CreatedBy, r.UpdatedBy)
	}

	mapUserIdName, _ := s.store.GetUserNamesByIds(userIds)

	resp := make([]dto.RecommendResponse, len(recommends))
	for i, r := range recommends {
		resp[i] = dto.RecommendResponse{
			Id:            r.Id,
			Title:         r.Title,
			Position:      r.Position,
			ActiveFrom:    r.ActiveFrom.Format(time.RFC3339),
			CreatedAt:     r.CreatedAt.Format(time.RFC3339),
			UpdatedAt:     r.UpdatedAt.Format(time.RFC3339),
			CreatedByName: mapUserIdName[r.CreatedBy],
			UpdatedByName: mapUserIdName[r.UpdatedBy],
		}
		if r.ActiveTo != nil {
			resp[i].ActiveTo = r.ActiveTo.Format(time.RFC3339)
		}

		if r.Cover != "" {
			resp[i].Cover = s.minio.GetFileUrl(s.config.FileStorage.RecommendFolder, r.Cover)
		}
	}
	config.BuildListResponse(ctx, &common.Pagination{Page: req.Page, PageSize: req.PageSize, Total: int(total)}, resp)
}

// @Summary Update recommend by Id
// @Description Update an existing recommend
// @Tags recommends
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Bearer authorization token"
// @Param title formData string true "Title"
// @Param position formData int false "Position"
// @Param active_from formData int false "Active From"
// @Param active_to formData int false "Active To"
// @Param cover formData file false "Cover "
// @Security BearerAuth
// @Success 200 {object} dto.ResponseMessage{data=dto.RecommendResponse} "Recommend updated successfully"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/recommend [put]
func (s *Server) UpdateRecommendById(ctx *gin.Context) {
	var req dto.RecommendUpdateRequest
	if err := req.Bind(ctx); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	file, _ := ctx.FormFile("cover")
	if file != nil {
		fileName, err := s.minio.SaveImage(file, s.config.FileStorage.RecommendFolder)
		if err != nil {
			config.BuildErrorResponse(ctx, err, nil)
			return
		}

		req.Cover = fileName
	}

	recommend, err := s.store.GetRecommendById(req.Id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	isUpdate := false
	if len(req.Title) > 0 {
		isUpdate = true
		recommend.Title = req.Title
	}

	if len(req.Cover) > 0 {
		isUpdate = true
		recommend.Cover = req.Cover
	}

	if req.Position > 0 {
		isUpdate = true
		recommend.Position = req.Position
	}
	if req.ActiveFrom > 0 {
		isUpdate = true
		recommend.ActiveFrom = config.FromLongValueToTime(req.ActiveFrom)
	}

	if req.ActiveTo > 0 {
		isUpdate = true
		recommend.ActiveTo = config.FromLongValueToTime(req.ActiveTo)
	}

	now := time.Now()
	var resp *dto.RecommendResponse
	if isUpdate {
		userId, err := s.GetUserIdFromContext(ctx)
		if err != nil {
			config.BuildErrorResponse(ctx, err, nil)
			return
		}
		recommend.UpdatedAt = &now
		recommend.UpdatedBy = userId

		err = s.store.UpdateRecommend(recommend)
		if err != nil {
			config.BuildErrorResponse(ctx, err, nil)
			return
		}
		resp = &dto.RecommendResponse{
			Id:         recommend.Id,
			Title:      recommend.Title,
			Cover:      recommend.Cover,
			Position:   recommend.Position,
			ActiveFrom: recommend.ActiveFrom.Format(time.RFC3339),
			ActiveTo:   recommend.ActiveTo.Format(time.RFC3339),
			CreatedAt:  recommend.CreatedAt.Format(time.RFC3339),
			UpdatedAt:  recommend.UpdatedAt.Format(time.RFC3339),
		}
		mapUserIdName, err := s.store.GetUserNamesByIds([]int64{recommend.CreatedBy, recommend.UpdatedBy})
		if err == nil && len(mapUserIdName) > 0 {
			resp.CreatedByName = mapUserIdName[recommend.CreatedBy]
			resp.UpdatedByName = mapUserIdName[recommend.UpdatedBy]
		}
	}

	if resp != nil {
		if resp.Cover != "" {
			resp.Cover = s.minio.GetFileUrl(s.config.FileStorage.RecommendFolder, resp.Cover)
		}

		config.BuildSuccessResponse(ctx, resp)
		return
	}
	config.BuildErrorResponse(ctx, errors.New("error update recommend"), nil)
}

// @Summary Delete recommend by Id
// @Description Delete a recommend by Id
// @Tags recommends
// @Accept json
// @Produce json
// @Param     Authorization header string true "Bearer authorization token"
// @Param id path int true "Recommend ID"
// @Security     BearerAuth
// @Success 200 {object} dto.ResponseMessage "Recommend successfully deleted"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/recommend/{id} [delete]
func (s *Server) DeleteRecommendById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		config.BuildErrorResponse(ctx, errors.New("invalid id"), nil)
		return
	}

	recommend, err := s.store.DeleteRecommendById(id)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	resp := dto.RecommendResponse{
		Id:         recommend.Id,
		Title:      recommend.Title,
		Cover:      recommend.Cover,
		Position:   recommend.Position,
		ActiveFrom: recommend.ActiveFrom.Format(time.RFC3339),
		ActiveTo:   recommend.ActiveTo.Format(time.RFC3339),
		CreatedAt:  recommend.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  recommend.UpdatedAt.Format(time.RFC3339),
	}

	mapUserIdName, err := s.store.GetUserNamesByIds([]int64{recommend.CreatedBy, recommend.UpdatedBy})
	if err == nil && len(mapUserIdName) > 0 {
		resp.CreatedByName = mapUserIdName[recommend.CreatedBy]
		resp.UpdatedByName = mapUserIdName[recommend.UpdatedBy]
	}
	config.BuildSuccessResponse(ctx, resp)
}

// @Summary Create recommend comic
// @Description Create a new recommend comic
// @Tags recommends
// @Accept json
// @Produce json
// @Param     Authorization header string true "Bearer authorization token"
// @Param recommend body dto.RecommendComicRequest true "Recommend comic create Request"
// @Security     BearerAuth
// @Success 200 {object} dto.ResponseMessage{data=dto.RecommendComicResponse} "Recommend comic created successfully"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/recommend/comic [post]
func (s *Server) CreateRecommendComic(ctx *gin.Context) {
	var req dto.RecommendComicRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	recommendComicModel := model.RecommendComicModel{
		RecommendId: req.RecommendId,
		ComicId:     req.ComicId,
	}

	err := s.store.CreateRecommendComic(&recommendComicModel)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	resp := dto.RecommendComicResponse{
		Id:          recommendComicModel.Id,
		RecommendId: recommendComicModel.RecommendId,
		ComicId:     recommendComicModel.ComicId,
	}

	config.BuildSuccessResponse(ctx, resp)
}

// @Summary Delete recommend comic
// @Description Delete a recommend comic
// @Tags recommends
// @Accept json
// @Produce json
// @Param     Authorization header string true "Bearer authorization token"
// @Param recommend body dto.RecommendComicRequest true "Recommend comic create Request"
// @Security     BearerAuth
// @Success 200 {object} dto.ResponseMessage{data=dto.RecommendComicResponse} "Recommend comic deleted successfully"
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/recommend/comic [delete]
func (s *Server) DeleteRecommendComic(ctx *gin.Context) {
	var req dto.RecommendComicRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	recommendComic, err := s.store.DeleteRecommendComicById(req.ComicId, req.RecommendId)
	if err != nil {
		config.BuildErrorResponse(ctx, err, nil)
		return
	}

	resp := dto.RecommendComicResponse{
		Id:          recommendComic.Id,
		RecommendId: recommendComic.RecommendId,
		ComicId:     recommendComic.ComicId,
	}
	config.BuildSuccessResponse(ctx, resp)
}
