package repository

import "be-client/internal/infra/database"

type UserComicHistoryRepository interface {
}

type userComicHistoryRepository struct {
	db database.DBInterface
}

func NewUserComicHistoryRepository(db database.DBInterface) UserComicHistoryRepository {
	return &userComicHistoryRepository{db: db}
}
