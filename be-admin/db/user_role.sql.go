package db

import (
	"context"
	"pkg-common/model"
)

func (q *Queries) CreateUserRole(userRole *model.UserRoleModel) error {
	return q.db.WithContext(context.Background()).Create(userRole).Error
}
