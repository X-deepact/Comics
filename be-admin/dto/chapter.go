package dto

type ChapterRequest struct {
	ComicId int64  `json:"comic_id" binding:"required"`
	Name    string `json:"name"`
	Cover   bool   `json:"cover"`
	Number  int    `json:"number"`
	Active  bool   `json:"active"`
}

type ChapterResponse struct {
	Id            int64  `json:"id"`
	ComicId       int64  `json:"comic_id"`
	Name          string `json:"name"`
	Cover         bool   `json:"cover"`
	Number        int    `json:"number"`
	Active        bool   `json:"active"`
	CreatedAt     string `json:"created_at"`
	CreatedBy     string `json:"created_by"`
	UpdatedAt     string `json:"updated_at"`
	UpdatedBy     string `json:"updated_by"`
	CreatedByUser string `json:"created_by_user"`
	UpdatedByUser string `json:"updated_by_user"`
}

type ChapterUpdateRequest struct {
	Id int64 `json:"id" binding:"required"`
	ChapterRequest
}

type ChapterListRequest struct {
	ListRequest
	ComicId int64  `form:"comic_id" json:"comic_id"`
	Order   string `form:"order" json:"order"`
}
