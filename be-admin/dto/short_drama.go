package dto

type ShortDramaRequest struct {
	ReleaseDate       string             `json:"release_date"`
	Thumb             string             `json:"thumb"`
	DramaTranslations []DramaTranslation `json:"translations"`
	Genres            []int64            `json:"genres"`
}

type UpdateShortDramaRequest struct {
	Id int64 `json:"id" binding:"required"`
	ShortDramaRequest
}

type DramaTranslation struct {
	Language    string `json:"language"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GenreForShortDramaResponse struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Position int64  `json:"position"`
}

type ShortDrama struct {
	Id            int64  `json:"id"`
	ReleaseDate   string `json:"release_date"`
	Thumb         string `json:"thumb"`
	Active        bool   `json:"active"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	CreatedByName string `json:"created_by_name"`
	UpdatedByName string `json:"updated_by_name"`
}

type ShortDramaResponse struct {
	ShortDrama
	Translations []DramaTranslation           `json:"translations"`
	Genres       []GenreForShortDramaResponse `json:"genres"`
}

type ShortDramaListRequest struct {
	ListRequest
	SortBy string `form:"sort_by" json:"sort_by"`
	Sort   string `form:"sort" json:"sort"`
}
