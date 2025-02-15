package db

import (
	"comics-admin/dto"
	"context"
	"pkg-common/model"
)

// CreateChapterItem creates a new chapter item
func (q *Queries) CreateChapterItem(item *model.ChapterItemModel) error {
	return q.db.WithContext(context.Background()).Create(item).Error
}

// GetChapterItem gets chapter item by ID
func (q *Queries) GetChapterItem(id int64) (*dto.ChapterItemResponse, error) {
	var item dto.ChapterItemResponse
	if err := q.db.WithContext(context.Background()).
		Table("chapter_items").
		Select("chapter_items.*, uc.username AS created_by_name, up.username AS updated_by_name").
		Joins("LEFT JOIN users uc ON uc.id = chapter_items.created_by").
		Joins("LEFT JOIN users up ON up.id = chapter_items.updated_by").
		Where("chapter_items.id = ?", id).
		First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

// ListChapterItems gets list of items for a chapter
func (q *Queries) ListChapterItems(req dto.ChapterItemListRequest) ([]dto.ChapterItemResponse, int64, error) {
	var items []dto.ChapterItemResponse
	var total int64

	query := q.db.WithContext(context.Background()).
		Table("chapter_items").
		Where("chapter_id = ?", req.ChapterId).
		Order("page ASC")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Select("chapter_items.*, uc.username AS created_by_name, up.username AS updated_by_name").
		Joins("LEFT JOIN users uc ON uc.id = chapter_items.created_by").
		Joins("LEFT JOIN users up ON up.id = chapter_items.updated_by").
		Limit(req.PageSize).
		Offset((req.Page - 1) * req.PageSize).
		Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

// UpdateChapterItem updates a chapter item
func (q *Queries) UpdateChapterItem(item *model.ChapterItemModel) error {
	return q.db.WithContext(context.Background()).Model(item).Updates(item).Error
}

// DeleteChapterItem deletes a chapter item
func (q *Queries) DeleteChapterItem(id int64) error {
	return q.db.WithContext(context.Background()).Delete(&model.ChapterItemModel{Id: id}).Error
}
