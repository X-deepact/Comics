package db

import (
	"context"
	"pkg-common/model"
)

func (r *Queries) CreateSubtitle(ctx context.Context, input *model.SubtitleModel) error {
	return r.db.WithContext(ctx).Create(input).Error
}

func (r *Queries) CreateSubtitles(ctx context.Context, input []*model.SubtitleModel) error {
	return r.db.WithContext(ctx).Create(&input).Error
}

func (r *Queries) GetSubtitleByEpisodeId(ctx context.Context, episodeId int64) ([]*model.SubtitleModel, error) {
	var subtitle []*model.SubtitleModel
	if err := r.db.WithContext(ctx).Where("episode_id = ?", episodeId).Find(&subtitle).Error; err != nil {
		return nil, err
	}
	return subtitle, nil
}

func (r *Queries) GetSubtitleByEpisodeIds(ctx context.Context, epsilonIds []int64) (map[int64][]*model.SubtitleModel, error) {
	var subtitle []*model.SubtitleModel
	if err := r.db.WithContext(ctx).Where("episode_id in (?)", epsilonIds).Find(&subtitle).Error; err != nil {
		return nil, err
	}
	result := make(map[int64][]*model.SubtitleModel)
	for _, subtitle := range subtitle {
		if _, ok := result[subtitle.EpisodeId]; !ok {
			result[subtitle.EpisodeId] = []*model.SubtitleModel{}
		}
		result[subtitle.EpisodeId] = append(result[subtitle.EpisodeId], subtitle)
	}
	return result, nil
}

func (r *Queries) DeleteSubtitleByEpisodeId(ctx context.Context, episodeId int64) error {
	return r.db.WithContext(ctx).Where("episode_id = ?", episodeId).Delete(&model.SubtitleModel{}).Error
}

func (r *Queries) UpdateSubtitleUrlByEpisodeId(ctx context.Context, epsilonId, updatedBy int64, input []*model.SubtitleModel) error {
	return r.db.WithContext(ctx).Model(&model.SubtitleModel{}).
		Save(input).Error
}
