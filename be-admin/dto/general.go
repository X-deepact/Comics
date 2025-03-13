package dto

type TierModel struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
}

type GeneralGenreRequest struct {
	Name     string `form:"name"`
	Language string `form:"language"`
}

type GeneralGenreResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GeneralAuthorRequest struct {
	Name string `form:"name"`
}

type GeneralAuthorResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GeneralComicRequest struct {
	Name             string `form:"name"`
	Language         string `form:"lang"`
	OriginalLanguage string `form:"original_language"`
	Audience         string `form:"audience"`
	Author           int64  `form:"author"`
	Genre            int64  `form:"genre"`
}

type GeneralComicResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
