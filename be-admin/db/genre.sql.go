package db

import (
	"comics-admin/dto"
	"context"
	"pkg-common/model"
)

func (q *Queries) CreateGenre(genre *model.GenreModel) error {
	return q.db.WithContext(context.Background()).Create(genre).Error
}

func (q *Queries) GetGenre(id int64) (*dto.GenreResponse, error) {
	var genre dto.GenreResponse
	if err := q.db.WithContext(context.Background()).
		Table("genres").
		Select("genres.*, uc.username AS created_by_name, up.username AS updated_by_name").
		Joins("LEFT JOIN users uc ON uc.id = genres.created_by").
		Joins("LEFT JOIN users up ON up.id = genres.updated_by").
		Where("genres.id = ?", id).
		First(&genre).Error; err != nil {
		return nil, err
	}
	return &genre, nil
}

func (q *Queries) GetGenres(req dto.GenreListRequest) ([]dto.GenreResponse, int64, error) {
	var genres []dto.GenreResponse
	var total int64
	var query = q.db.WithContext(context.Background()).Table("genres")

	if req.Name != "" {
		query.Where("genres.name LIKE ?", "%"+req.Name+"%")
	}

	if req.Language != "" {
		query.Where("genres.lang = ?", req.Language)
	}

	query = query.Select("genres.*, uc.username AS created_by_name, up.username AS updated_by_name").
		Joins("LEFT JOIN users uc ON uc.id = genres.created_by").
		Joins("LEFT JOIN users up ON up.id = genres.updated_by")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("genres.position").
		Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).
		Find(&genres).Error; err != nil {
		return nil, 0, err
	}

	return genres, total, nil
}

func (q *Queries) UpdateGenre(genre *model.GenreModel) error {
	return q.db.WithContext(context.Background()).Model(genre).Updates(genre).Error
}

func (q *Queries) DeleteGenre(id int64) error {
	return q.db.WithContext(context.Background()).Delete(&model.GenreModel{Id: id}).Error
}
