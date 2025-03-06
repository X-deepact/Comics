package dto

type GenreCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Position int    `json:"position" binding:"required"`
	Language string `json:"language" binding:"required"`
}

type GenreResponse struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Position      int    `json:"position"`
	Lang          string `json:"lang"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	CreatedByName string `json:"created_by_name"`
	UpdatedByName string `json:"updated_by_name"`
}

type GenreListRequest struct {
	ListRequest
	Name     string `form:"name" json:"name"`
	Language string `form:"language" json:"language"`
	SortBy   string `form:"sort_by" json:"sort_by"`
	Sort     string `form:"sort" json:"sort"`
}

type GenreUpdateRequest struct {
	ID       int64  `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Position int    `json:"position" binding:"required"`
	Language string `json:"language" binding:"required"`
}
