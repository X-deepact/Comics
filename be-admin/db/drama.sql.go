package db

import (
	"comics-admin/dto"
	"context"
	"pkg-common/model"

	"gorm.io/gorm"
)

func (q *Queries) CreateDrama(drama *model.ShortDramaModel) error {
	return q.db.WithContext(context.Background()).Create(drama).Error
}

func (q *Queries) CreateDramaTranslations(dramaTranslations []*model.DramaTranslationModel) error {
	return q.db.WithContext(context.Background()).Create(dramaTranslations).Error
}

func (q *Queries) CreateDramaGenres(dramaGenres []*model.DramaGenreModel) error {
	return q.db.WithContext(context.Background()).Create(dramaGenres).Error
}

func (q *Queries) GetDrama(id int64) (*model.ShortDramaModel, error) {
	var drama *model.ShortDramaModel
	if err := q.db.WithContext(context.Background()).Model(&model.ShortDramaModel{}).Where("id = ?", id).First(&drama).Error; err != nil {
		return nil, err
	}
	return drama, nil

}

func (q *Queries) GetDramaTranslations(dramaId int64) ([]dto.DramaTranslation, error) {
	var translations []dto.DramaTranslation
	if err := q.db.WithContext(context.Background()).
		Table("drama_translations").
		Where("drama_id = ?", dramaId).
		Find(&translations).Error; err != nil {
		return nil, err
	}
	return translations, nil
}

func (q *Queries) GetDramaGenres(dramaId int64) ([]dto.GenreForShortDramaResponse, error) {
	var genres []dto.GenreForShortDramaResponse
	if err := q.db.WithContext(context.Background()).
		Table("drama_genres").
		Select("genres_for_short_drama.*").
		Joins("LEFT JOIN genres_for_short_drama ON genres_for_short_drama.id = drama_genres.genre_id").
		Where("drama_id = ?", dramaId).
		Find(&genres).Error; err != nil {
		return nil, err
	}
	return genres, nil
}

func (q *Queries) GetDramas(req dto.ShortDramaListRequest) ([]*model.ShortDramaModel, int64, error) {
	var dramas []*model.ShortDramaModel
	var total int64
	query := q.db.WithContext(context.Background()).Model(&model.ShortDramaModel{})
	if req.SortBy != "" {
		order := req.SortBy + " " + req.Sort
		query = query.Order(order)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.PageSize
	query = query.Offset(offset).Limit(req.PageSize)

	if err := query.Find(&dramas).Error; err != nil {
		return nil, 0, err
	}

	return dramas, total, nil
}

func (q *Queries) UpdateDrama(drama *model.ShortDramaModel) error {
	return q.db.WithContext(context.Background()).Model(drama).Updates(drama).Error
}

func (q *Queries) DeleteDramaTranslations(dramaId int64) error {
	return q.db.WithContext(context.Background()).Delete(&model.DramaTranslationModel{}, "drama_id = ?", dramaId).Error
}

func (q *Queries) DeleteDramaGenres(dramaId int64) error {
	return q.db.WithContext(context.Background()).Delete(&model.DramaGenreModel{}, "drama_id = ?", dramaId).Error
}

func (q *Queries) DeleteDrama(id int64) error {
	return q.db.WithContext(context.Background()).Delete(&model.ShortDramaModel{}, id).Error
}

func (q *Queries) ActiveDrama(id int64, adminId int64) error {
	activeDrama := map[string]interface{}{
		"active":     gorm.Expr("NOT active"),
		"updated_by": adminId,
	}

	return q.db.WithContext(context.Background()).Model(&model.ShortDramaModel{}).Where("id = ?", id).Updates(activeDrama).Error
}
