package db

import (
	"comics-admin/dto"
	"context"
)

func (q *Queries) GetTiers() ([]*dto.TierModel, error) {
	var tier []*dto.TierModel
	if err := q.db.WithContext(context.Background()).Table("tiers").
		Select("id, code").
		Find(&tier).Error; err != nil {
		return nil, err
	}
	return tier, nil
}
