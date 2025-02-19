package db

import (
	"context"
	"pkg-common/model"
)

func (r *Queries) CreateRecomend(input *model.RecommendManagerModel) error {
	return r.db.WithContext(context.Background()).Create(input).Error
}

func (r *Queries) GetRecommends(limit, offset int) ([]*model.RecommendManagerModel, int64, error) {
	var recommends []*model.RecommendManagerModel
	var total int64

	if err := r.db.WithContext(context.Background()).Limit(limit).Offset(offset).Find(&recommends).Count(&total).Error; err != nil {
		return recommends, total, err
	}

	return recommends, total, nil
}

func (r *Queries) GetRecommendById(id int64) (*model.RecommendManagerModel, error) {
	var recommend model.RecommendManagerModel
	if err := r.db.WithContext(context.Background()).Where("id = ?", id).First(&recommend).Error; err != nil {
		return nil, err
	}
	return &recommend, nil
}

func (r *Queries) DeleteRecomendById(id int64) (*model.RecommendManagerModel, error) {
	recommend := &model.RecommendManagerModel{}
	err := r.db.WithContext(context.Background()).Delete(recommend, "id = ?", id).Error
	if err != nil {
		return recommend, err
	}
	return recommend, nil
}

func (r *Queries) UpdateRecomend(input *model.RecommendManagerModel) error {
	return r.db.WithContext(context.Background()).Save(input).Error
}
