package response

type SubtitleResponse struct {
	EpisodeId   int64  `json:"episode_id"`
	Language    string `json:"language"`
	SubtitleUrl string `json:"subtitle_url"`
}
