package response

import "time"

type ComicResponse struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Language    string `json:"language"`
	Audience    string `json:"audience"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
	Active      bool   `json:"active"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type ComicWithChapterResponse struct {
	Id          int64                  `json:"id"`
	Name        string                 `json:"name"`
	Code        string                 `json:"code"`
	Language    string                 `json:"language"`
	Audience    string                 `json:"audience"`
	Description string                 `json:"description"`
	Cover       string                 `json:"cover"`
	Active      bool                   `json:"active"`
	CreatedAt   int64                  `json:"created_at"`
	UpdatedAt   int64                  `json:"updated_at"`
	Chapters    []ComicChapterResponse `json:"chapters"`
}

type ComicWithChapterAndAuthorResponse struct {
	Id          int64                      `json:"id"`
	Name        string                     `json:"name"`
	Code        string                     `json:"code"`
	Language    string                     `json:"language"`
	Audience    string                     `json:"audience"`
	Description string                     `json:"description"`
	Cover       string                     `json:"cover"`
	Active      bool                       `json:"active"`
	CreatedAt   *time.Time                 `json:"created_at"`
	UpdatedAt   *time.Time                 `json:"updated_at"`
	Chapters    []ComicChapterTimeResponse `json:"chapters"`
	Authors     []ComicAuthorTimeResponse  `json:"authors"`
}

type ComicDetailResponse struct {
	Id          int64                  `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Code        string                 `json:"code,omitempty"`
	Language    string                 `json:"language,omitempty"`
	Audience    string                 `json:"audience,omitempty"`
	Description string                 `json:"description,omitempty"`
	Cover       string                 `json:"cover,omitempty"`
	Active      bool                   `json:"active,omitempty"`
	Status      string                 `json:"status,omitempty"`
	CreatedAt   int64                  `json:"created_at,omitempty"`
	UpdatedAt   int64                  `json:"updated_at,omitempty"`
	Chapters    []ComicChapterResponse `json:"chapters,omitempty"`
	Genres      []ComicGenreResponse   `json:"genres,omitempty"`
	Authors     []ComicAuthorResponse  `json:"authors,omitempty"`
}

type ComicAuthorResponse struct {
	Id        int64  `json:"id,omitempty"`
	Biography string `json:"biography,omitempty"`
	Name      string `json:"name,omitempty"`
	BirthDay  int64  `json:"birth_day,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type ComicGenreResponse struct {
	Id        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Position  int    `json:"position,omitempty"`
	Lang      string `json:"lang,omitempty"`
	CreatedBy int64  `json:"created_by,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedBy int64  `json:"updated_by,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type ComicChapterResponse struct {
	Id        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	IsCover   bool   `json:"is_cover,omitempty"`
	Number    int64  `json:"number,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}

type ComicRecent struct {
	Id                 int64                      `json:"id"`
	Name               string                     `json:"name"`
	Code               string                     `json:"code"`
	Language           string                     `json:"language"`
	Audience           string                     `json:"audience"`
	Description        string                     `json:"description"`
	Cover              string                     `json:"cover"`
	Active             bool                       `json:"active"`
	LastChapterCreDate *time.Time                 `json:"latest_date"`
	Chapters           []ComicChapterTimeResponse `json:"chapters"`
}

type ComicChapterTimeResponse struct {
	Id        int64      `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	IsCover   bool       `json:"is_cover,omitempty"`
	Number    int64      `json:"number,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

type ComicAuthorTimeResponse struct {
	Id        int64      `json:"id,omitempty"`
	Biography string     `json:"biography,omitempty"`
	Name      string     `json:"name,omitempty"`
	BirthDay  *time.Time `json:"birth_day,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ComicWithChapterTimeResponse struct {
	Id          int64                      `json:"id"`
	Name        string                     `json:"name"`
	Code        string                     `json:"code"`
	Language    string                     `json:"language"`
	Audience    string                     `json:"audience"`
	Description string                     `json:"description"`
	Cover       string                     `json:"cover"`
	Active      bool                       `json:"active"`
	CreatedAt   *time.Time                 `json:"created_at"`
	UpdatedAt   *time.Time                 `json:"updated_at"`
	Chapters    []ComicChapterTimeResponse `json:"chapters"`
}

type ComicAndDramaSearchResponse struct {
	Id                  int64                      `json:"id"`
	Type                string                     `json:"type"`
	Name                string                     `json:"name"`
	Code                string                     `json:"code,omitempty"`
	Language            string                     `json:"language"`
	Audience            string                     `json:"audience,omitempty"`
	Description         string                     `json:"description"`
	Cover               string                     `json:"cover,omitempty"`
	Active              bool                       `json:"active"`
	CreatedAt           *time.Time                 `json:"created_at"`
	UpdatedAt           *time.Time                 `json:"updated_at"`
	ReleaseDate         *time.Time                 `json:"release_date,omitempty"`
	LatestEpisodeNumber int64                      `json:"latest_episode"`
	Thumb               string                     `json:"thumb,omitempty"`
	Chapters            []ComicChapterTimeResponse `json:"chapters,omitempty"`
	Authors             []ComicAuthorTimeResponse  `json:"authors,omitempty"`
}
