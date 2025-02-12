package db

import (
	"context"
	"pkg-common/model"
)

func (q *Queries) GetTiers(code string) (*model.TierModel, error) {
	var tier model.TierModel
	if err := q.db.WithContext(context.Background()).Where("code = ?", code).First(&tier).Error; err != nil {
		return nil, err
	}
	return &tier, nil
}
