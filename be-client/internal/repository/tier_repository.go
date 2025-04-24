package repository

import (
	"be-client/internal/infra/database"
	"context"
	"pkg-common/model"
)

type TierRepository interface {
	GetTierById(ctx context.Context, id int64) (*model.TierModel, error)
}

type tierRepository struct {
	db database.DBInterface
}

func NewTierRepository(db database.DBInterface) TierRepository {
	return &tierRepository{db: db}
}

func (r *tierRepository) GetTierById(ctx context.Context, id int64) (*model.TierModel, error) {
	req := model.TierModel{}
	err := r.db.Ctx(ctx).Where("id = ?", id).First(&req)
	return &req, err
}
