package dto

type GenreDramaCreateRequest struct {
	Name         string                    `json:"name" binding:"required"`
	Position     int64                     `json:"position" binding:"required"`
	Translations []GenreTranslationRequest `json:"translations" binding:"required,min=1"`
}

type GenreDramaUpdateRequest struct {
	ID           int64                     `json:"id" binding:"required"`
	Name         string                    `json:"name" binding:"required"`
	Position     int64                     `json:"position" binding:"required"`
	Translations []GenreTranslationRequest `json:"translations" binding:"required,min=1"`
}

type GenreDramaListRequest struct {
	Page     int    `form:"page" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,min=10,max=50"`
	Name     string `form:"name"`
	Language string `form:"language"`
	SortBy   string `form:"sort_by"`
	Sort     string `form:"sort"`
}

type GenreTranslationRequest struct {
	Language       string `json:"language" binding:"required"`
	TranslatedName string `json:"translated_name" binding:"required"`
}

type GenreDramaResponse struct {
	ID           int64                      `json:"id"`
	Name         string                     `json:"name"`
	Position     int64                      `json:"position"`
	CreatedAt    string                     `json:"created_at"`
	UpdatedAt    string                     `json:"updated_at"`
	CreatedBy    int64                      `json:"created_by"`
	UpdatedBy    int64                      `json:"updated_by"`
	Translations []GenreTranslationResponse `json:"translations,omitempty"`
}

type GenreTranslationResponse struct {
	ID             int64  `json:"id"`
	GenreId        int64  `json:"genre_id"`
	Language       string `json:"language"`
	TranslatedName string `json:"translated_name"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	CreatedBy      int64  `json:"created_by"`
	UpdatedBy      int64  `json:"updated_by"`
}
