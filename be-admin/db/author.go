package db

import (
	"context"
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

func (q *Queries) GetAuthors(limit, offset int) ([]*model.AuthorModel, int64, error) {
	var authors []*model.AuthorModel
	var total int64

	if err := q.db.WithContext(context.Background()).Limit(limit).Offset(offset).Find(&authors).Count(&total).Error; err != nil {
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
	err := q.db.WithContext(context.Background()).Delete(author, "id = ?", id).Error
	if err != nil {
		return author, err
	}
	return author, nil
}
