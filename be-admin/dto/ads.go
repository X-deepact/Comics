package dto

import "time"

type AdsCreateRequest struct {
	Title      string     `json:"title" binding:"required"`
	Image      string     `json:"image"`
	ComicID    int64      `json:"comic_id"`
	Type       string     `json:"type" binding:"required,oneof=internal external"`
	Status     string     `json:"status" binding:"required,oneof=active inactive"`
	DirectURL  string     `json:"direct_url"`
	ActiveFrom *time.Time `json:"active_from"`
	ActiveTo   *time.Time `json:"active_to"`
}

type AdsUpdateRequest struct {
	ID         int64      `json:"id" binding:"required"`
	Title      string     `json:"title"`
	Image      string     `json:"image"`
	ComicID    int64      `json:"comic_id"`
	Type       string     `json:"type" binding:"oneof=internal external"`
	Status     string     `json:"status" binding:"oneof=active inactive"`
	DirectURL  string     `json:"direct_url"`
	ActiveFrom *time.Time `json:"active_from"`
	ActiveTo   *time.Time `json:"active_to"`
}

type AdsResponse struct {
	ID            int64      `json:"id"`
	Title         string     `json:"title"`
	Image         string     `json:"image"`
	ComicID       int64      `json:"comic_id"`
	Type          string     `json:"type"`
	Status        string     `json:"status"`
	DirectURL     string     `json:"direct_url"`
	ActiveFrom    *time.Time `json:"active_from"`
	ActiveTo      *time.Time `json:"active_to"`
	CreatedAt     time.Time  `json:"created_at"`
	CreatedByName string     `json:"created_by_name"`
	UpdatedAt     time.Time  `json:"updated_at"`
	UpdatedByName string     `json:"updated_by_name"`
}

type AdsListRequest struct {
	Page     int    `form:"page" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,min=1,max=100"`
	Title    string `form:"title"`
	Type     string `form:"type" binding:"omitempty,oneof=internal external"`
	Status   string `form:"status" binding:"omitempty,oneof=active inactive"`
	SortBy   string `form:"sort_by"`
	Sort     string `form:"sort"`
}

type AdsUpdateStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=active inactive"`
}
