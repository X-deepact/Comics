package request

type GetPopularComicsRequest struct {
	Genre            int    `query:"genre_id"`
	Progress         string `query:"progress"`
	Audience         string `query:"audience"`
	Language         string `query:"language"`
	OriginalLanguage string `query:"original_language"`
	Page             int    `query:"page"`
	PageSize         int    `query:"page_size"`
}

type GetComicSearchRequest struct {
	SearchKeyword string `query:"search_keyword"`
	Language      string `query:"language"`
	Page          int    `query:"page"`
	PageSize      int    `query:"page_size"`
}

type GetComicRecentRequest struct {
	Language string `query:"language"`
	Page     int    `query:"page"`
	PageSize int    `query:"page_size"`
}
