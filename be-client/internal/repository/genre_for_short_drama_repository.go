package repository

import (
	"be-client/internal/dto/request"
	"be-client/internal/infra/database"
	"be-client/util"
	"context"
	"pkg-common/model"
)

type GenreForShortDramaRepository interface {
	GetAllGenreForShortDrama(ctx context.Context, req request.GetGenresForShortDramaRequest) (*[]map[string]interface{}, error)
	GetAllGenreByDramaId(ctx context.Context, dramaId int64, req request.GetShortDramaDetailRequest) (*[]map[string]interface{}, error)
}

type genreForShortDramaRepository struct {
	db database.DBInterface
}

func NewGenreForShortDramaRepository(db database.DBInterface) GenreForShortDramaRepository {
	return &genreForShortDramaRepository{db}
}

func (r *genreForShortDramaRepository) GetAllGenreForShortDrama(ctx context.Context, req request.GetGenresForShortDramaRequest) (*[]map[string]interface{}, error) {
	var results []map[string]interface{}
	query := r.db.Ctx(ctx).Model(&model.GenreForShortDramaModel{}).
		Select(`genres_for_short_drama.*,gt.language,gt.translated_name`).
		Joins("join genre_translations gt ON gt.genre_id = genres_for_short_drama.id")

	if req.Language != "" {
		query = query.Where("gt.language = ?", req.Language)
	} else {
		languge := util.GetDefaultLanguage()
		query = query.Where("gt.language = ?", languge)
	}

	err := query.Find(&results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}

func (r *genreForShortDramaRepository) GetAllGenreByDramaId(ctx context.Context, dramaId int64, req request.GetShortDramaDetailRequest) (*[]map[string]interface{}, error) {
	var results []map[string]interface{}
	query := r.db.Ctx(ctx).Model(&model.GenreForShortDramaModel{}).
		Select(`genres_for_short_drama.*,gt.language,gt.translated_name`).
		Joins("join genre_translations gt ON gt.genre_id = genres_for_short_drama.id").
		Joins(`join drama_genres dg ON dg.genre_id = genres_for_short_drama.id and dg.drama_id = ?`, dramaId).
		Order(`genres_for_short_drama.position`)

	if req.Language != "" {
		query = query.Where("gt.language = ?", req.Language)
	} else {
		languge := util.GetDefaultLanguage()
		query = query.Where("gt.language = ?", languge)
	}

	err := query.Find(&results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}
