package db

import (
	"comics-admin/dto"
	"pkg-common/model"

	"gorm.io/gorm"
)

type Store interface {
	Querier
	CreateGenre(genre *model.GenreModel) error
	GetGenre(id int64) (*dto.GenreResponse, error)
	GetGenres(req dto.GenreListRequest) ([]dto.GenreResponse, int64, error)
	UpdateGenre(genre *model.GenreModel) error
	DeleteGenre(id int64) error

	// Methods for genre drama
	CreateGenreDrama(genre *model.GenreForShortDramaModel) error
	CreateGenreTranslation(translation *model.GenreTranslationModel) error
	CreateGenreTranslations(translations []*model.GenreTranslationModel) error
	GetGenreDrama(id int64) (*dto.GenreDramaResponse, error)
	UpdateGenreDrama(genre *model.GenreForShortDramaModel) error
	DeleteGenreTranslations(genreId int64) error
	DeleteGenreDrama(id int64) error
	GetGenreDramas(req dto.GenreDramaListRequest) ([]dto.GenreDramaResponse, int64, error)
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
