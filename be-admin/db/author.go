package db

import (
	"comics-admin/dto"
	"context"
	"fmt"
	"pkg-common/model"
)

func (q *Queries) CreateAuthor(author *model.AuthorModel) error {
	return q.db.WithContext(context.Background()).Create(author).Error
}

func (q *Queries) GetAuthorById(id int64) (*model.AuthorModel, error) {
	var author model.AuthorModel
	if err := q.db.WithContext(context.Background()).Where("id = ?", id).First(&author).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func (q *Queries) GetAuthors(req dto.RequestQueryFilter) ([]*model.AuthorModel, int64, error) {
	var authors []*model.AuthorModel
	var total int64

	query := q.db.WithContext(context.Background()).Model(&model.AuthorModel{})
	if req.Name != "" {
		query = query.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if err := query.
		Count(&total).
		Order(fmt.Sprintf("%s %s", req.SortBy, req.Sort)).
		Limit(req.PageSize).
		Offset((req.Page - 1) * req.PageSize).
		Find(&authors).Error; err != nil {
		return authors, total, err
	}

	return authors, total, nil
}

func (q *Queries) UpdateAuthor(author *model.AuthorModel) (*model.AuthorModel, error) {
	err := q.db.WithContext(context.Background()).Save(author).Error
	if err != nil {
		return author, err
	}

	return author, nil
}

func (q *Queries) DeleteAuthorById(id int64) (*model.AuthorModel, error) {
	author := &model.AuthorModel{}
	if err := q.db.WithContext(context.Background()).First(author, id).Error; err != nil {
		return nil, err
	}

	if err := q.db.WithContext(context.Background()).Delete(author).Error; err != nil {
		return nil, err
	}

	return author, nil
}

func (q *Queries) GetAuthorsOfAComic(comicId int64) ([]model.AuthorModel, error) {
	var authors []model.AuthorModel
	query := q.db.WithContext(context.Background()).Table("authors").
		Joins("JOIN comic_authors ON authors.id = comic_authors.author_id").
		Where("comic_authors.comic_id = ?", comicId)

	if err := query.Find(&authors).Error; err != nil {
		return nil, err
	}

	return authors, nil
}
