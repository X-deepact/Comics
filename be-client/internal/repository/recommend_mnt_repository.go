package repository

import (
	"be-client/internal/infra/database"
	"context"
	"pkg-common/model"
	"time"
)

type RecommendManagerRepository interface {
	GetAllRecommend(ctx context.Context) (*[]model.RecommendManagerModel, []int64, error)
}

type recommendManagerRepository struct {
	db database.DBInterface
}

func NewRecommendManagerRepository(db database.DBInterface) RecommendManagerRepository {
	return &recommendManagerRepository{db: db}
}

func (r *recommendManagerRepository) GetAllRecommend(ctx context.Context) (*[]model.RecommendManagerModel, []int64, error) {
	var recommends []model.RecommendManagerModel
	query := r.db.Ctx(ctx).
		Where("active_from < ? AND (active_to IS NULL OR active_to > ?)", time.Now(), time.Now())

	err := query.Order("position ASC").Find(&recommends)
	if err != nil {
		return nil, nil, err
	}

	var rmIds []int64
	for _, recommend := range recommends {
		if rmId := recommend.Id; rmId > 0 {
			rmIds = append(rmIds, rmId)
		}
	}

	return &recommends, rmIds, nil
}
