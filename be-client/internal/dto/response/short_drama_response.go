package response

import "time"

type ShortDramaSearchResponse struct {
	Id          int64      `json:"id"`
	Name        string     `json:"name"`
	ReleaseDate *time.Time `json:"release_date"`
	Thumb       string     `json:"thumb"`
	Language    string     `json:"language"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type ShortDramaByClassificationResponse struct {
	Id                  int64      `json:"id"`
	Name                string     `json:"name"`
	TranslatedName      string     `json:"translated_name"`
	ReleaseDate         *time.Time `json:"release_date"`
	Thumb               string     `json:"thumb"`
	Language            string     `json:"language"`
	Description         string     `json:"description"`
	LatestEpisodeNumber int64      `json:"latest_episode"`
	CreatedAt           *time.Time `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
}

type ShortDramaDetailResponse struct {
	Id             int64                         `json:"id"`
	Name           string                        `json:"name"`
	TranslatedName string                        `json:"translated_name"`
	ReleaseDate    *time.Time                    `json:"release_date"`
	Thumb          string                        `json:"thumb"`
	Language       string                        `json:"language"`
	Description    string                        `json:"description"`
	CreatedAt      *time.Time                    `json:"created_at"`
	UpdatedAt      *time.Time                    `json:"updated_at"`
	Episodes       []EpisodeResponse             `json:"episodes"`
	Genres         []GenresForShortDramaResponse `json:"genres"`
}

type EpisodeResponse struct {
	Id        int64      `json:"id"`
	Number    int64      `json:"number"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
