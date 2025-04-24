package response

import "time"

type EpisodeGetResponse struct {
	Video     string     `json:"video"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
