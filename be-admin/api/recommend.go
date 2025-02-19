package api

import (
	"comics-admin/dto"
	config "comics-admin/util"
	"fmt"
	"net/http"
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
}

// @Summary Create recommend
// @Description Create a new recommend
// @Tags recommends
// @Accept json
// @Produce json
// @Param recommend body dto.RecommendCreateRequest true "Recommend Create Request"
// @Security     BearerAuth
// @Success 200 {object} dto.RecommendResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/recommend [post]
func (s *Server) CreateRecommend(ctx *gin.Context) {
	var req dto.RecommendCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("user not authenticated")))
		return
	}
	idInt64, err := strconv.ParseInt(userID.(string), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
		CreatedBy:  idInt64,
		UpdatedBy:  idInt64,
	}

	err = s.store.CreateRecomend(&recommendModel)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
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
		CreatedBy:  recommendModel.CreatedBy,
		UpdatedAt:  recommendModel.UpdatedAt.Format(time.RFC3339),
		UpdatedBy:  recommendModel.UpdatedBy,
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get recommend by Id
// @Description Get a recommend by Id
// @Tags recommends
// @Accept json
// @Produce json
// @Param id path int true "Recommend ID"
// @Security     BearerAuth
// @Success 200 {object} dto.RecommendResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/recommend/{id} [get]
func (s *Server) GetRecommendById(ctx *gin.Context) {
	var req dto.ListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponseMessage("invalid id"))
		return
	}

	recommend, err := s.store.GetRecommendById(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
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
		CreatedBy:  recommend.CreatedBy,
		UpdatedAt:  recommend.UpdatedAt.Format(time.RFC3339),
		UpdatedBy:  recommend.UpdatedBy,
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary List recommends
// @Description List all recommends
// @Tags recommends
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Security     BearerAuth
// @Success 200 {object} []dto.RecommendResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/recommend/list [get]
func (s *Server) GetRecommends(ctx *gin.Context) {
	var req dto.ListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	offset := (req.Page - 1) * req.PageSize
	recommend, total, err := s.store.GetRecommends(req.PageSize, offset)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	resp := make([]dto.RecommendResponse, len(recommend))
	for i, r := range recommend {
		resp[i] = dto.RecommendResponse{
			Id:         r.Id,
			Title:      r.Title,
			Cover:      r.Cover,
			Position:   r.Position,
			ActiveFrom: r.ActiveFrom.Format(time.RFC3339),
			ActiveTo:   r.ActiveTo.Format(time.RFC3339),
			CreatedAt:  r.CreatedAt.Format(time.RFC3339),
			CreatedBy:  r.CreatedBy,
			UpdatedAt:  r.UpdatedAt.Format(time.RFC3339),
			UpdatedBy:  r.UpdatedBy,
		}
	}
	ListResponse(ctx, req.Page+1, req.PageSize, int(total), resp)
}

// @Summary Update recommend by Id
// @Description Update an existing recommend
// @Tags recommends
// @Accept json
// @Produce json
// @Param id path int true "Recommend ID"
// @Param recommend body dto.RecommendUpdateRequest true "Recommend Update Request"
// @Security     BearerAuth
// @Success 200 {object} dto.RecommendResponse
// @Failure 400 {object} dto.ResponseMessage "Invalid request"
// @Failure 500 {object} dto.ResponseMessage "Internal server error"
// @Router /api/recommend/{id} [put]
func (s *Server) UpdateRecommendById(ctx *gin.Context) {
	var req dto.RecommendUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponseMessage("invalid id"))
		return
	}

	recommend, err := s.store.GetRecommendById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
	if len(req.ActiveFrom) > 0 {
		isUpdate = true
		activeFrom, err := config.ConvertStringToDate(req.ActiveFrom)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		recommend.ActiveFrom = &activeFrom
	}

	if len(req.ActiveTo) > 0 {
		isUpdate = true
		activeTo, err := config.ConvertStringToDate(req.ActiveTo)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		recommend.ActiveTo = &activeTo
	}

	var resp *dto.RecommendResponse
	if isUpdate {
		err := s.store.UpdateRecomend(recommend)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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
			CreatedBy:  recommend.CreatedBy,
			UpdatedAt:  recommend.UpdatedAt.Format(time.RFC3339),
			UpdatedBy:  recommend.UpdatedBy,
		}
	}

	if resp != nil {
		ctx.JSON(http.StatusOK, resp)
		return
	}
	ctx.JSON(http.StatusBadRequest, errorResponseMessage("error update recommend"))
}

func (s *Server) DeleteRecommendById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponseMessage("invalid id"))
		return
	}

	recommend, err := s.store.DeleteRecomendById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponseMessage(err.Error()))
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
		CreatedBy:  recommend.CreatedBy,
		UpdatedAt:  recommend.UpdatedAt.Format(time.RFC3339),
		UpdatedBy:  recommend.UpdatedBy,
	}
	ctx.JSON(http.StatusOK, resp)
}
