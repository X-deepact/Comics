package dto

import "time"

// ChapterItemCreateRequest represents request to create new chapter item
type ChapterItemCreateRequest struct {
	ChapterId  int64      `json:"chapter_id" binding:"required"`
	Page       int64      `json:"page" binding:"required"`
	Image      string     `json:"image" binding:"required"`
	Active     bool       `json:"active"`
	ActiveFrom *time.Time `json:"active_from"`
}

// ChapterItemUpdateRequest represents request to update chapter item
type ChapterItemUpdateRequest struct {
	Id         int64      `json:"id" binding:"required"`
	Page       int64      `json:"page"`
	Image      string     `json:"image"`
	Active     bool       `json:"active"`
	ActiveFrom *time.Time `json:"active_from"`
}

// ChapterItemResponse represents chapter item data in responses
type ChapterItemResponse struct {
	Id            int64      `json:"id"`
	ChapterId     int64      `json:"chapter_id"`
	Page          int64      `json:"page"`
	Image         string     `json:"image"`
	Active        bool       `json:"active"`
	ActiveFrom    *time.Time `json:"active_from"`
	CreatedAt     time.Time  `json:"created_at"`
	CreatedByName string     `json:"created_by_name"`
	UpdatedAt     time.Time  `json:"updated_at"`
	UpdatedByName string     `json:"updated_by_name"`
}

// ChapterItemListRequest represents request to list chapter items
type ChapterItemListRequest struct {
	ChapterId int64 `form:"chapter_id" binding:"required"`
	Page      int   `form:"page" binding:"required,min=1"`
	PageSize  int   `form:"page_size" binding:"required,min=1,max=100"`
}
