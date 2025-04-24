package repository

import (
	"be-client/internal/dto/request"
	"be-client/internal/infra/database"
	"context"
	"pkg-common/model"
)

type DramaTranslationRepository interface {
	SearchDramaFromName(ctx context.Context, req request.GetComicSearchRequest) ([]int64, error)
}

type dramaTranslationRepository struct {
	db database.DBInterface
}

func NewDramaTranslationRepository(db database.DBInterface) DramaTranslationRepository {
	return &dramaTranslationRepository{db}
}

func (r *dramaTranslationRepository) SearchDramaFromName(ctx context.Context, req request.GetComicSearchRequest) ([]int64, error) {
	var results []int64

	queryDrama := r.db.Ctx(ctx).
		Model(&model.ShortDramaModel{}).
		Select(`d_tran.drama_id`).
		Joins("JOIN drama_translations d_tran ON d_tran.drama_id = short_dramas.id").
		Where("short_dramas.active = ?", true).
		Group("d_tran.drama_id")

	if req.SearchKeyword != "" {
		queryDrama = queryDrama.Where("(d_tran.title LIKE ?)", "%"+req.SearchKeyword+"%")
	}

	err := queryDrama.Find(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
