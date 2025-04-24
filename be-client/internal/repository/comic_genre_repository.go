package repository

import "be-client/internal/infra/database"

type ComicGenreRepository interface {
}

type comicGenreRepository struct {
	db database.DBInterface
}

func NewComicGenreRepository(db database.DBInterface) ComicGenreRepository {
	return &comicGenreRepository{db: db}
}
