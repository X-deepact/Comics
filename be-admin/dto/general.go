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
