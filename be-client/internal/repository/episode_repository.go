package repository

import (
	"be-client/internal/infra/database"
	"context"
	"pkg-common/model"
)

type EpisodeRepository interface {
	GetById(ctx context.Context, epId int64) (*model.EpisodeModel, error)
	GetLatestEpNumByDranaIds(ctx context.Context, dramaIds []int64) (*[]map[string]interface{}, error)
	GetAllByDramaId(ctx context.Context, dramaId int64) (*[]model.EpisodeModel, error)
}

type episodeRepository struct {
	db database.DBInterface
}

func NewEpisodeRepository(db database.DBInterface) EpisodeRepository {
	return &episodeRepository{db}
}

func (r *episodeRepository) GetById(ctx context.Context, epId int64) (*model.EpisodeModel, error) {
	var drama model.EpisodeModel

	query := r.db.Ctx(ctx).Model(&model.EpisodeModel{}).Where(`active = ?`, true).Where(`id = ?`, epId)

	err := query.Find(&drama)
	if err != nil {
		return nil, err
	}

	return &drama, nil
}

func (r *episodeRepository) GetLatestEpNumByDranaIds(ctx context.Context, dramaIds []int64) (*[]map[string]interface{}, error) {
	var results []map[string]interface{}
	limit := 1

	query := `SELECT * FROM (
			SELECT drama_id, number,
				   ROW_NUMBER() OVER (PARTITION BY drama_id ORDER BY created_at DESC, number DESC) AS rn
			FROM episodes
			WHERE active = ? AND drama_id IN (?)
		) ch_filtered
		WHERE rn <= ?`
	err := r.db.Ctx(ctx).
		Raw(query, true, dramaIds, limit).Find(&results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}

func (r *episodeRepository) GetAllByDramaId(ctx context.Context, dramaId int64) (*[]model.EpisodeModel, error) {
	var drama []model.EpisodeModel

	query := r.db.Ctx(ctx).Model(&model.EpisodeModel{}).
		Where(`active = ?`, true).
		Where(`drama_id = ?`, dramaId).
		Order(`number`)

	err := query.Find(&drama)
	if err != nil {
		return nil, err
	}

	return &drama, nil
}
