package db

import (
	"context"
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

func (q *Queries) ListComics(limit, offset int) ([]model.ComicModel, int64, error) {
	var comics []model.ComicModel
	var total int64

	if err := q.db.WithContext(context.Background()).Limit(limit).Offset(offset).Find(&comics).Error; err != nil {
		return nil, 0, err
	}

	if err := q.db.WithContext(context.Background()).Model(&model.ComicModel{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return comics, total, nil
}

func (q *Queries) UpdateComic(comic *model.ComicModel) error {
	return q.db.WithContext(context.Background()).Model(comic).Updates(comic).Error
}
