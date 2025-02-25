package db

import (
	"comics-admin/dto"
	"context"
	"gorm.io/gorm"
	"pkg-common/model"
)

func (q *Queries) CreateChapter(chapter *model.ChapterModel) error {
	return q.db.WithContext(context.Background()).Create(chapter).Error
}

func (q *Queries) GetChapter(id int64) (*dto.ChapterResponse, error) {
	var chapter dto.ChapterResponse
	if err := q.db.WithContext(context.Background()).
		Table("chapters").
		Select("chapters.*, uc.username AS created_by_name, up.username AS updated_by_name").
		Joins("LEFT JOIN users uc ON uc.id = chapters.created_by").
		Joins("LEFT JOIN users up ON up.id = chapters.updated_by").
		Where("chapters.id = ?", id).
		First(&chapter).Error; err != nil {
		return nil, err
	}
	return &chapter, nil
}

func (q *Queries) ListChapters(req dto.ChapterListRequest) ([]dto.ChapterResponse, int64, error) {
	var chapters []dto.ChapterResponse
	var total int64

	query := q.db.WithContext(context.Background()).Table("chapters")

	query = query.Select("chapters.*, uc.username AS created_by_name, up.username AS updated_by_name").
		Joins("LEFT JOIN users uc ON uc.id = chapters.created_by").
		Joins("LEFT JOIN users up ON up.id = chapters.updated_by")

	if req.ComicId != 0 {
		query = query.Where("comic_id = ?", req.ComicId)
	}

	if req.SortBy != "" {
		order := req.SortBy + " " + req.Sort
		query = query.Order(order)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.PageSize
	query = query.Offset(offset).Limit(req.PageSize)

	if err := query.Find(&chapters).Error; err != nil {
		return nil, 0, err
	}

	return chapters, total, nil
}

func (q *Queries) UpdateChapter(chapter *model.ChapterModel) error {
	return q.db.WithContext(context.Background()).Model(chapter).Updates(chapter).Error
}

func (q *Queries) DeleteChapter(id int64) error {
	return q.db.WithContext(context.Background()).Delete(&model.ChapterModel{}, id).Error
}

func (q *Queries) ActiveChapter(id int64, adminId int64) error {
	activeChapter := map[string]interface{}{
		"active":     gorm.Expr("NOT active"),
		"updated_by": adminId,
	}

	return q.db.WithContext(context.Background()).Model(&model.ChapterModel{}).Where("id = ?", id).Updates(activeChapter).Error
}
