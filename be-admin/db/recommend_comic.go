package db

import (
	"context"
	"pkg-common/model"
)

func (r *Queries) CreateRecommendComic(input *model.RecommendComicModel) error {
	return r.db.WithContext(context.Background()).Create(input).Error
}

func (r *Queries) DeleteRecommendComicById(comicId, recommendId int64) (*model.RecommendComicModel, error) {
	resp := &model.RecommendComicModel{}

	if err := r.db.WithContext(context.Background()).Where("comic_id = ? AND recommend_id = ?", comicId, recommendId).First(resp).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(context.Background()).Delete(resp).Error; err != nil {
		return nil, err
	}

	return resp, nil
}
