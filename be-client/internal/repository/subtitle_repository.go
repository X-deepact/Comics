package repository

import (
	"be-client/internal/infra/database"
	"context"
	"pkg-common/model"
)

type SubtitleRepository interface {
	GetVTTsByEpisodeId(ctx context.Context, episodeId int64) ([]model.SubtitleModel, error)
}

type subtitleRepository struct {
	db database.DBInterface
}

func NewSubtitleRepository(db database.DBInterface) SubtitleRepository {
	return &subtitleRepository{db}
}

func (r *subtitleRepository) GetVTTsByEpisodeId(ctx context.Context, episodeId int64) ([]model.SubtitleModel, error) {
	var subtitles []model.SubtitleModel
	err := r.db.Ctx(ctx).Where("episode_id = ?", episodeId).Find(&subtitles)
	if err != nil {
		return nil, err
	}
	return subtitles, nil
}
