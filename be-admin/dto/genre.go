package dto

type GenreCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Position int    `json:"position" binding:"required"`
	Language string `json:"language" binding:"required"`
}

type GenreResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Position      int    `json:"position"`
	Language      string `json:"language"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	CreatedByName string `json:"created_by_name"`
	UpdatedByName string `json:"updated_by_name"`
}

type GenreListRequest struct {
	Page     int    `form:"page" json:"page" binding:"required,min=1"`
	PageSize int    `form:"page_size" json:"page_size" binding:"required,min=1,max=100"`
	Name     string `form:"name" json:"name"`
	Language string `form:"language" json:"language"`
}

type GenreUpdateRequest struct {
	ID       int64  `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Position int    `json:"position" binding:"required"`
	Language string `json:"language" binding:"required"`
}
