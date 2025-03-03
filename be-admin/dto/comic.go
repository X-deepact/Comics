package dto

import (
	"pkg-common/model"
	"time"
)

type ComicRequest struct {
	Name        string  `form:"name" binding:"required"`
	Code        string  `form:"code"`
	Description string  `form:"description"`
	Lang        string  `form:"lang"`
	Audience    string  `form:"audience"`
	CreatedBy   int64   `form:"created_by"`
	UpdatedBy   int64   `form:"updated_by"`
	Status      string  `form:"status"`
	Authors     []int64 `form:"authors"`
	Genres      []int64 `form:"genres"`
	Cover       string  `json:"cover"`
}

type ComicListRequest struct {
	ListRequest
	Query       string `form:"q" json:"q"`
	SortBy      string `form:"sort_by" json:"sort_by"`
	Sort        string `form:"sort" json:"sort"`
	Active      bool   `form:"active" json:"active"`
	ActiveValue bool
	Lang        string `form:"lang" json:"lang"`
	Audience    string `form:"audience" json:"audience"`
	Author      int64  `form:"author" json:"author"`
	Genre       int64  `form:"genre" json:"genre"`
}

type ComicResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
	Lang        string `json:"lang"`
	Audience    string `json:"audience"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	CreatedBy   int64  `json:"created_by"`
	UpdatedBy   int64  `json:"updated_by"`
}

type ComicReturn struct {
	ComicResponse
	CreatedByUser UserDetailDto       `json:"created_by_user"`
	UpdatedByUser UserDetailDto       `json:"updated_by_user"`
	Genres        []GenreResponse     `json:"genres"`
	Authors       []model.AuthorModel `json:"authors"`
}

type ComicUpdateRequest struct {
	ID int64 `form:"id" binding:"required"`
	ComicRequest
}

type AuthorModelReturn struct {
	Id        int64      `gorm:"primaryKey;column:id"`
	Biography string     `gorm:"column:biography"`
	Name      string     `gorm:"column:name"`
	BirthDay  *time.Time `gorm:"column:birth_day"`
	CreatedBy int64      `gorm:"column:created_by"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedBy int64      `gorm:"column:updated_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

type ComicModelReturn struct {
	ComicResponse
	CreatedByUser UserDetailDto       `json:"created_by_user"`
	UpdatedByUser UserDetailDto       `json:"updated_by_user"`
	Genres        []GenreResponse     `json:"genres"`
	Authors       []AuthorModelReturn `json:"authors"`
}
