package response

import "time"

type RecommendComicResponse struct {
	Id         int64                          `json:"id"`
	Title      string                         `json:"title"`
	Cover      string                         `json:"cover"`
	Position   string                         `json:"position"`
	ActiveFrom *time.Time                     `json:"active_from"`
	ActiveTo   *time.Time                     `json:"active_to"` // nullable
	CreatedAt  *time.Time                     `json:"created_at"`
	UpdatedAt  *time.Time                     `json:"updated_at"`
	Comics     []ComicWithChapterTimeResponse `json:"comics"`
}
