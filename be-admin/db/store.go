package db

import (
	"gorm.io/gorm"
)

type Store interface {
	Querier
}

type SQLStore struct {
	*Queries
	DB *gorm.DB
}

func NewStore(db *gorm.DB) *SQLStore {
	return &SQLStore{
		Queries: New(db),
		DB:      db,
	}
}
