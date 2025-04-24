package repository

import "be-client/internal/infra/database"

type RoleRepository interface {
}

type roleRepository struct {
	db database.DBInterface
}

func NewRoleRepository(db database.DBInterface) RoleRepository {
	return &roleRepository{db: db}
}
