package response

type GenreResponse struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Position int    `json:"position"`
}
