package response

import "time"

type GenresForShortDramaResponse struct {
	Id             int64      `json:"id"`
	Name           string     `json:"name"`
	TranslatedName string     `json:"translated_name"`
	Position       int64      `json:"position"`
	Language       string     `json:"language"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
