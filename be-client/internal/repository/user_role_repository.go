package repository

import "be-client/internal/infra/database"

type UserRoleRepository interface {
}

type userRoleRepository struct {
	db database.DBInterface
}

func NewUserRoleRepository(db database.DBInterface) UserRoleRepository {
	return &userRoleRepository{db: db}
}
