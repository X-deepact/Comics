package repository

import (
	"be-client/internal/infra/database"
	"context"
	"pkg-common/model"
)

type ComicAuthorRepository interface {
	GetComicAuthorsByComicId(ctx context.Context, comicId int64) ([]*model.ComicAuthorModel, error)
	GetComicAndAuthorsByComicId(ctx context.Context, comicIDs []int64) (*[]map[string]interface{}, error)
}

type comicAuthorRepository struct {
	db database.DBInterface
}

func NewComicAuthorRepository(db database.DBInterface) ComicAuthorRepository {
	return &comicAuthorRepository{db: db}
}

func (r *comicAuthorRepository) GetComicAuthorsByComicId(ctx context.Context, comicId int64) ([]*model.ComicAuthorModel, error) {
	var comicAuthors []*model.ComicAuthorModel
	err := r.db.Ctx(ctx).
		Where("comic_id = ?", comicId).
		Find(&comicAuthors)
	if err != nil {
		return nil, err
	}
	return comicAuthors, nil
}

func (r *comicAuthorRepository) GetComicAndAuthorsByComicId(ctx context.Context, comicIDs []int64) (*[]map[string]interface{}, error) {
	var comicAuthors []map[string]interface{}
	err := r.db.Ctx(ctx).Model(&model.ComicAuthorModel{}).
		Select(`comic_authors.comic_id as cm_id,comic_authors.author_id as aut_id,aut.name as aut_name, 
			aut.id as aut_id, aut.biography as aut_biography, aut.birth_day as aut_birthday, 
			aut.created_at as aut_created_at, aut.updated_at as updated_at`).
		Joins("JOIN authors aut ON aut.id = comic_authors.author_id").
		Where(`comic_authors.comic_id IN (?)`, comicIDs).
		Find(&comicAuthors)
	if err != nil {
		return nil, err
	}
	return &comicAuthors, nil
}
