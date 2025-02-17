package dto

type ComicRequest struct {
	Name        string `json:"title" binding:"required"`
	Code        string `json:"code"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
	Language    string `json:"language"`
	Audience    string `json:"audience"`
}

type ComicListRequest struct {
	ListRequest
	Query    string `form:"q" json:"q"`
	SortBy   string `form:"sort_by" json:"sort_by"`
	Sort     string `form:"sort" json:"sort"`
	Active   bool   `form:"active" json:"active"`
	Language string `form:"language" json:"language"`
	Audience string `form:"audience" json:"audience"`
}

type ComicResponse struct {
	ID            int64  `json:"id"`
	Name          string `json:"title"`
	Code          string `json:"code"`
	Cover         string `json:"cover"`
	Description   string `json:"description"`
	Active        bool   `json:"active"`
	Language      string `json:"language"`
	Audience      string `json:"audience"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	CreatedBy     int64  `json:"created_by"`
	UpdatedBy     int64  `json:"updated_by"`
	CreatedByName string `json:"created_by_name"`
	UpdatedByName string `json:"updated_by_name"`
}

type ComicUpdateRequest struct {
	ID int64 `json:"id" binding:"required"`
	ComicRequest
}
