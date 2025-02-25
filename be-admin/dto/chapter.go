package dto

type ChapterRequest struct {
	ComicId int64  `json:"comic_id" binding:"required"`
	Name    string `json:"name"`
	Cover   bool   `json:"cover"`
	Number  int    `json:"number"`
}

type ChapterResponse struct {
	Id            int64  `json:"id"`
	ComicId       int64  `json:"comic_id"`
	Name          string `json:"name"`
	Cover         bool   `json:"cover"`
	Number        int    `json:"number"`
	Active        bool   `json:"active"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	CreatedByName string `json:"created_by_name"`
	UpdatedByName string `json:"updated_by_name"`
}

type ChapterUpdateRequest struct {
	Id int64 `json:"id" binding:"required"`
	ChapterRequest
}

type ChapterListRequest struct {
	ListRequest
	ComicId int64  `form:"comic_id" json:"comic_id"`
	SortBy  string `form:"sort_by" json:"sort_by"`
	Sort    string `form:"sort" json:"sort"`
}
