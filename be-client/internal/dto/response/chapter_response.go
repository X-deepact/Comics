package response

type ChapterDetailResponse struct {
	Id           int64                 `json:"id,omitempty"`
	Name         string                `json:"name,omitempty"`
	ComicId      int64                 `json:"comic_id,omitempty"`
	ComicName    string                `json:"comic_name,omitempty"`
	Cover        bool                  `json:"cover,omitempty"`
	Number       int                   `json:"number,omitempty"`
	Active       bool                  `json:"active,omitempty"`
	ActiveFrom   int64                 `json:"active_from,omitempty"`
	CreatedAt    int64                 `json:"created_at,omitempty"`
	ChapterItems []ChapterItemResponse `json:"items,omitempty"`
	NextChapter  *int                  `json:"next_chapter"`
	PrevChapter  *int                  `json:"prev_chapter"`
}

type ChapterItemResponse struct {
	Page       int64  `json:"page,omitempty"`
	Active     bool   `json:"active,omitempty"`
	Image      string `json:"image,omitempty"`
	ActiveFrom int64  `json:"active_from,omitempty"`
	CreatedAt  int64  `json:"created_at,omitempty"`
}
