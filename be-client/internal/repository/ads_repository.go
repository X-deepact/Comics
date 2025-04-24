package repository

import (
	"be-client/internal/dto/request"
	"be-client/internal/infra/database"
	"context"
	"pkg-common/model"
	"time"
)

type AdsRepository interface {
	GetAdsComics(ctx context.Context, request request.GetAdsComicRequest) (*[]map[string]interface{}, error)
}

type adsRepository struct {
	db database.DBInterface
}

func NewAdsRepository(db database.DBInterface) AdsRepository {
	return &adsRepository{db}
}

func (r *adsRepository) GetAdsComics(ctx context.Context, req request.GetAdsComicRequest) (*[]map[string]interface{}, error) {
	var results []map[string]interface{}
	query := r.db.Ctx(ctx).
		Model(&model.AdModel{}).
		Select(`ads.id AS ads_id, ads.title AS ads_title, ads.image AS ads_image, 
			ads.type AS ads_type, ads.direct_url AS ads_direct_url, 
			ads.active_from AS ads_active_from, ads.active_to AS ads_active_to, 
			cm.id AS comic_id, cm.name AS comic_name, 
			cm.code AS comic_code, cm.lang AS comic_language, 
			cm.audience AS comic_audience, cm.description AS comic_description, 
			cm.cover AS comic_cover, cm.active AS comic_active`).
		Joins("LEFT JOIN comics cm ON cm.id = ads.comic_id").
		Where("ads.active_from < ? AND (ads.active_to IS NULL OR ads.active_to > ?)", time.Now(), time.Now())

	if req.Language != "" {
		query = query.Where("cm.lang = ? or ads.type = 'external' ", req.Language)
	}

	query = query.Order("ads.updated_at DESC")

	limit := 5
	query = query.Limit(limit)

	err := query.Find(&results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}
