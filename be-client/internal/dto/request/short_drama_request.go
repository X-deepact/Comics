package request

type GetShortDramaByClassificationRequest struct {
	Genre       int    `query:"genre_id"`
	ReleaseYear string `query:"release_year"`
	Language    string `query:"language"`
	Page        int    `query:"page"`
	PageSize    int    `query:"page_size"`
}

type GetShortDramaDetailRequest struct {
	Language string `query:"language"`
}
