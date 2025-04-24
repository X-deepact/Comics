package repository

import (
	"be-client/internal/dto/request"
	"be-client/internal/infra/database"
	"be-client/util"
	"context"
	"pkg-common/model"
)

type ShortDramaRepository interface {
	GetShortDramaByClassification(ctx context.Context, request request.GetShortDramaByClassificationRequest) (*[]map[string]interface{}, int64, []int64, error)
	GetShortDramaById(ctx context.Context, id int64, req request.GetShortDramaDetailRequest) (*map[string]interface{}, error)
}

type shortDramaRepository struct {
	db database.DBInterface
}

func NewShortDramaRepository(db database.DBInterface) ShortDramaRepository {
	return &shortDramaRepository{db}
}

func (r *shortDramaRepository) GetShortDramaByClassification(ctx context.Context, req request.GetShortDramaByClassificationRequest) (*[]map[string]interface{}, int64, []int64, error) {
	var results []map[string]interface{}
	var total int64

	limit := req.PageSize
	page := req.Page

	if page <= 0 {
		page = 1
	}

	query := r.db.Ctx(ctx).
		Model(&model.ShortDramaModel{}).
		Select(`short_dramas.id AS d_id, short_dramas.release_date AS d_release_date,short_dramas.thumb as d_thumb,
			short_dramas.created_at AS d_created_at, short_dramas.updated_at AS d_updated_at, short_dramas.name AS d_name,
			d_tran.language AS d_tran_language, d_tran.title AS d_tran_title,
			d_tran.description AS d_tran_description`).
		Joins("JOIN drama_translations d_tran ON d_tran.drama_id = short_dramas.id").
		Where(`short_dramas.active = ?`, true)

	if req.Language != "" {
		query = query.Where("d_tran.language = ?", req.Language)
	} else {
		languge := util.GetDefaultLanguage()
		query = query.Where("d_tran.language = ?", languge)
	}

	if req.Genre > 0 {
		query = query.Joins(`JOIN drama_genres dg ON dg.drama_id = short_dramas.id`).
			Where(`dg.genre_id = ?`, req.Genre)
	}

	if req.ReleaseYear != "" {
		query = query.Where(`YEAR(short_dramas.release_date) = ?`, req.ReleaseYear)
	}

	if err := query.Count(&total); err != nil {
		return nil, 0, nil, err
	}

	if limit > 0 {
		offset := (page - 1) * limit
		query = query.Limit(limit).Offset(offset)
	}

	query = query.Order("short_dramas.updated_at DESC, d_tran.updated_at DESC")

	err := query.Find(&results)
	if err != nil {
		return nil, 0, nil, err
	}

	var dramaIDs []int64
	for _, drama := range results {
		if dramaID := util.ToInt64(drama["d_id"]); dramaID > 0 {
			dramaIDs = append(dramaIDs, dramaID)
		}
	}

	return &results, total, dramaIDs, nil
}

func (r *shortDramaRepository) GetShortDramaById(ctx context.Context, id int64, req request.GetShortDramaDetailRequest) (*map[string]interface{}, error) {
	var results map[string]interface{}

	query := r.db.Ctx(ctx).
		Model(&model.ShortDramaModel{}).
		Select(`short_dramas.id AS d_id, short_dramas.release_date AS d_release_date,short_dramas.thumb as d_thumb,
			short_dramas.created_at AS d_created_at, short_dramas.updated_at AS d_updated_at, short_dramas.name AS d_name,
			d_tran.language AS d_tran_language, d_tran.title AS d_tran_title,
			d_tran.description AS d_tran_description`).
		Joins("JOIN drama_translations d_tran ON d_tran.drama_id = short_dramas.id").
		Where(`short_dramas.id = ?`, id).
		Where(`short_dramas.active = ?`, true).
		Limit(1)

	if req.Language != "" {
		query = query.Where("d_tran.language = ?", req.Language)
	} else {
		languge := util.GetDefaultLanguage()
		query = query.Where("d_tran.language = ?", languge)
	}

	err := query.Find(&results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}
