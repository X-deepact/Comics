package repository

import (
	"be-client/internal/infra/database"
	"context"
	"pkg-common/model"

	"gorm.io/gorm"
)

type GenreRepository interface {
	GetAllGenres(isHomePage bool, lang string) (*[]model.GenreModel, error)
	GetGenresByComicId(ctx context.Context, id int64) ([]*model.GenreModel, error)
}

type genreRepository struct {
	db database.DBInterface
}

func NewGenreRepository(db database.DBInterface) GenreRepository {
	return &genreRepository{db}
}

func (r *genreRepository) GetAllGenres(isHomePage bool, lang string) (*[]model.GenreModel, error) {
	var genres []model.GenreModel
	query := r.db.Model(&model.GenreModel{})

	if isHomePage {
		query = query.Where("position BETWEEN ? AND ?", 1, 8)
	}

	if lang != "" {
		query = query.Where("lang = ?", lang)
	}

	if err := query.Find(&genres); err != nil {
		return nil, err
	}

	return &genres, nil
}

func (r *genreRepository) GetGenresByComicId(ctx context.Context, id int64) ([]*model.GenreModel, error) {
	var genres []*model.GenreModel
	err := r.db.Ctx(ctx).Table("genres g").
		Joins("JOIN comic_genres cg ON cg.genre_id = g.id").
		Joins("JOIN comics c ON c.id = cg.comic_id").
		Where("c.id = ?", id).Select("g.*").
		Find(&genres)
	if err != nil && err == gorm.ErrRecordNotFound {
		return genres, nil
	}
	return genres, nil
}
