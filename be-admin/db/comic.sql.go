package db

import (
	"comics-admin/dto"
	"context"
	"fmt"
	"pkg-common/model"
)

func (q *Queries) CreateComic(comic *model.ComicModel) error {
	return q.db.WithContext(context.Background()).Create(comic).Error
}

func (q *Queries) DeleteComic(id int64) error {
	return q.db.WithContext(context.Background()).Delete(&model.ComicModel{Id: id}).Error
}

func (q *Queries) GetComic(id int64) (*model.ComicModel, error) {
	var comic model.ComicModel
	if err := q.db.WithContext(context.Background()).Where("id = ?", id).First(&comic).Error; err != nil {
		return nil, err
	}
	return &comic, nil
}
func (q *Queries) ListComics(req dto.ComicListRequest) ([]dto.ComicResponse, int64, error) {
	var comics []dto.ComicResponse
	var total int64

	query := q.db.WithContext(context.Background()).Table("comics").
		Select("comics.*, created_by_user.username AS created_by_name, updated_by_user.username AS updated_by_name").
		Joins("LEFT JOIN users AS created_by_user ON comics.created_by = created_by_user.id").
		Joins("LEFT JOIN users AS updated_by_user ON comics.updated_by = updated_by_user.id")

	if req.Query != "" {
		query = query.Where("comics.name LIKE ? OR comics.code LIKE ?", "%"+req.Query+"%", "%"+req.Query+"%")
	}
	if req.Active {
		query = query.Where("comics.active = ?", req.Active)
	}
	if req.Language != "" {
		query = query.Where("comics.lang = ?", req.Language)
	}
	if req.Audience != "" {
		query = query.Where("comics.audience = ?", req.Audience)
	}

	if req.SortBy != "" {
		order := fmt.Sprintf("%s %s", req.SortBy, req.Sort)
		query = query.Order(order)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.PageSize
	query = query.Offset(offset).Limit(req.PageSize)

	if err := query.Find(&comics).Error; err != nil {
		return nil, 0, err
	}

	return comics, total, nil
}
func (q *Queries) UpdateComic(comic *model.ComicModel) error {
	return q.db.WithContext(context.Background()).Model(comic).Updates(comic).Error
}
