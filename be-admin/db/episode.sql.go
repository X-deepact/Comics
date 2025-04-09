package db

import (
	"comics-admin/dto"
	"context"
	"errors"
	"pkg-common/model"
)

func (q *Queries) CreateEpisode(episode *model.EpisodeModel) error {
	if episode == nil {
		return errors.New("invalid episode")
	}
	return q.db.WithContext(context.Background()).Create(episode).Error
}

func (q *Queries) GetEpisodeById(id int64) (*model.EpisodeModel, error) {
	var episode *model.EpisodeModel
	if err := q.db.WithContext(context.Background()).Where("id = ?", id).First(&episode).Error; err != nil {
		return nil, err
	}
	return episode, nil
}

func (q *Queries) ListEpisodes(req *dto.EpisodeListRequest) ([]*model.EpisodeModel, int64, error) {
	if req == nil {
		return nil, 0, errors.New("invalid request")
	}
	var episodes []*model.EpisodeModel
	var total int64
	query := q.db.WithContext(context.Background()).Model(&model.EpisodeModel{})
	if req.SortBy != "" {
		order := req.SortBy + " " + req.Sort
		query = query.Order(order)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.PageSize
	query = query.Offset(offset).Limit(req.PageSize)

	if err := query.Find(&episodes).Error; err != nil {
		return nil, 0, err
	}

	return episodes, total, nil
}

func (q *Queries) UpdateEpisode(episode *model.EpisodeModel) error {
	if episode == nil {
		return errors.New("invalid episode")
	}
	return q.db.WithContext(context.Background()).Model(episode).Save(episode).Error
}

func (q *Queries) DeleteEpisode(id int64) error {
	return q.db.WithContext(context.Background()).Delete(&model.EpisodeModel{}, id).Error
}
