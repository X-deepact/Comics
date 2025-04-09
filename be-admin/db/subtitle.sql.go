package db

import (
	"comics-admin/dto"
	"context"
	"pkg-common/model"
	"time"
)

func (r *Queries) CreateSubtitle(ctx context.Context, input *model.SubtitleModel) error {
	return r.db.WithContext(ctx).Create(input).Error
}

func (r *Queries) CreateSubtitles(ctx context.Context, input []*model.SubtitleModel) error {
	return r.db.WithContext(ctx).Create(&input).Error
}

func (r *Queries) GetSubtitleByEpisodeId(ctx context.Context, episodeId int64) ([]dto.Subtitle, error) {
	var subtitle []*model.SubtitleModel
	if err := r.db.WithContext(ctx).Where("episode_id = ?", episodeId).Find(&subtitle).Error; err != nil {
		return nil, err
	}
	result := make([]dto.Subtitle, len(subtitle))
	for i, subtitle := range subtitle {
		result[i] = dto.Subtitle{
			Language: subtitle.Language,
			Url:      subtitle.SubtitleUrl,
		}
	}
	return result, nil
}

func (r *Queries) GetSubtitleByEpisodeIds(ctx context.Context, epsilonIds []int64) (map[int64][]dto.Subtitle, error) {
	var subtitle []*model.SubtitleModel
	if err := r.db.WithContext(ctx).Where("episode_id in (?)", epsilonIds).Find(&subtitle).Error; err != nil {
		return nil, err
	}
	result := make(map[int64][]dto.Subtitle)
	for _, subtitle := range subtitle {
		if _, ok := result[subtitle.EpisodeId]; !ok {
			result[subtitle.EpisodeId] = make([]dto.Subtitle, 0)
		}
		result[subtitle.EpisodeId] = append(result[subtitle.EpisodeId], dto.Subtitle{
			Language: subtitle.Language,
			Url:      subtitle.SubtitleUrl,
		})
	}
	return result, nil
}

func (r *Queries) DeleteSubtitleByEpisodeId(ctx context.Context, episodeId int64) error {
	return r.db.WithContext(ctx).Where("episode_id = ?", episodeId).Delete(&model.SubtitleModel{}).Error
}

func (r *Queries) UpdateSubtitleUrlByEpisodeId(ctx context.Context, epsilonId, updatedBy int64, input []dto.Subtitle) error {
	now := time.Now()
	for i := range input {
		mapUpdate := map[string]interface{}{
			"subtitle_url": input[i].Url,
			"updated_at":   &now,
			"updated_by":   updatedBy,
		}
		subtitle := &model.SubtitleModel{
			Id:          epsilonId,
			SubtitleUrl: input[i].Url,
			UpdatedAt:   &now,
			UpdatedBy:   updatedBy,
		}
		if err := r.db.WithContext(ctx).Model(subtitle).Where("episode_id = ?", epsilonId).Updates(mapUpdate).Error; err != nil {
			return err
		}
	}
	return nil
}
