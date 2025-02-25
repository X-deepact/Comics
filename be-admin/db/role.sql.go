package db

import (
	"context"
	"pkg-common/model"
)

func (q *Queries) GetRole(name string) (*model.RoleModel, error) {
	var role model.RoleModel
	if err := q.db.WithContext(context.Background()).Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
