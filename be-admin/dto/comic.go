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

type ComicResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"title"`
	Code        string `json:"code"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	CreatedBy   int64  `json:"created_by"`
	UpdatedBy   int64  `json:"updated_by"`
}

type ComicUpdateRequest struct {
	ID          int64  `json:"id" binding:"required"`
	Name        string `json:"title"`
	Code        string `json:"code"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
	Language    string `json:"language"`
	Audience    string `json:"audience"`
	CreatedBy   int64  `json:"created_by"`
}
