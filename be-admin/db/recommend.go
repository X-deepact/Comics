package db

import (
	"comics-admin/dto"
	"context"
	"fmt"
	"pkg-common/model"
)

func (r *Queries) CreateRecommend(input *model.RecommendManagerModel) error {
	return r.db.WithContext(context.Background()).Create(input).Error
}

func (r *Queries) GetRecommends(req dto.RequestQueryFilter) ([]*model.RecommendManagerModel, int64, error) {
	var recommends []*model.RecommendManagerModel
	var total int64

	if err := r.db.WithContext(context.Background()).
		Order(fmt.Sprintf("%s %s", req.SortBy, req.Sort)).
		Limit(req.PageSize).
		Offset((req.Page - 1) * req.PageSize).
		Find(&recommends).
		Count(&total).Error; err != nil {
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

func (r *Queries) DeleteRecommendById(id int64) (*model.RecommendManagerModel, error) {
	resp := &model.RecommendManagerModel{}
	if err := r.db.WithContext(context.Background()).First(resp, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(context.Background()).Delete(resp).Error; err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Queries) UpdateRecommend(input *model.RecommendManagerModel) error {
	return r.db.WithContext(context.Background()).Save(input).Error
}
