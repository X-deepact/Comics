package repository

import (
	"be-client/internal/infra/database"
	"context"
	"log/slog"
	"pkg-common/model"
)

type AuthorRepository interface {
	GetAuthorsByComicIds(context.Context, []int64) ([]*model.AuthorModel, error)
}

type authorRepository struct {
	db database.DBInterface
}

func NewAuthorRepository(db database.DBInterface) AuthorRepository {
	return &authorRepository{db}
}

func (r *authorRepository) GetAuthorsByComicIds(ctx context.Context, ids []int64) ([]*model.AuthorModel, error) {
	slog.Info("get author  by comic ids ")
	var authors []*model.AuthorModel
	err := r.db.Ctx(ctx).Table("authors a").
		Joins("JOIN comic_authors ca ON ca.author_id = a.id").
		Joins("JOIN comics c ON c.id = ca.comic_id").
		Where("ca.comic_id IN ?", ids).
		Select("a.*").
		Find(&authors)
	return authors, err
}
