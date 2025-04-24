package repository

import (
	"be-client/internal/dto/request"
	"be-client/internal/infra/database"
	"be-client/util"
	"context"
)

type RecommendComicRepository interface {
	GetRecommendComics(ctx context.Context, req request.GetRecommendComicRequest, rmIds []int64) (*[]map[string]interface{}, []int64, error)
	GetRecommendComics1(ctx context.Context, req request.GetRecommendComicRequest, rmIds []int64) (*[]map[string]interface{}, []int64, error)
}

type recommendComicRepository struct {
	db database.DBInterface
}

func NewRecommendComicRepository(db database.DBInterface) RecommendComicRepository {
	return &recommendComicRepository{db: db}
}

func (r *recommendComicRepository) GetRecommendComics(ctx context.Context, req request.GetRecommendComicRequest, rmIds []int64) (*[]map[string]interface{}, []int64, error) {
	var recommends []map[string]interface{}
	var args []interface{}

	query := `
		WITH RankedComics AS (
			SELECT 
				rc.recommend_id AS rm_id,
				cm.id AS cm_id,
				cm.name AS cm_name,
				cm.code AS cm_code,
				cm.lang AS cm_language,
				cm.audience AS cm_audience,
				cm.description AS cm_description,
				cm.cover AS cm_cover,
				cm.active AS cm_active,
				rm.position AS rm_position,
				cm.status AS cm_status,
				cm.created_at AS cm_created_at,
				cm.updated_at AS cm_updated_at,
				ROW_NUMBER() OVER (
					PARTITION BY rm.position 
					ORDER BY 
						CASE WHEN cm.status = 'completed' THEN 1 ELSE 0 END DESC, 
						cm.updated_at DESC, 
						cm.id
				) AS rn
			FROM recommend_comics rc
			LEFT JOIN comics cm ON cm.id = rc.comic_id
			JOIN recommend_managers rm ON rm.id = rc.recommend_id
			WHERE 
				rc.recommend_id IN (?) 
				AND cm.id IS NOT NULL
				AND cm.active = TRUE
	`

	args = append(args, rmIds)

	if req.Language != "" {
		query += " AND cm.lang = ?"
		args = append(args, req.Language)
	}

	query += `
		)
		SELECT 
			rm_id, 
			cm_id, 
			cm_name, 
			cm_code, 
			cm_language, 
			cm_audience, 
			cm_description, 
			cm_cover, 
			cm_active, 
			rm_position, 
			cm_status, 
			cm_created_at,
			cm_updated_at
		FROM RankedComics 
		WHERE 
			(rm_position = 1 AND rn <= 5) OR 
			(rm_position = 2 AND rn <= 9) OR 
			(rm_position NOT IN (1,2) AND rn <= 10);
	`

	err := r.db.Ctx(ctx).
		Raw(query, args...).Find(&recommends)

	if err != nil {
		return nil, nil, err
	}

	var comicIDs []int64
	uniqueComicMap := make(map[int64]struct{})

	for _, recommend := range recommends {
		if cmID := util.ToInt64(recommend["cm_id"]); cmID > 0 {
			if _, exists := uniqueComicMap[cmID]; !exists {
				uniqueComicMap[cmID] = struct{}{}
				comicIDs = append(comicIDs, cmID)
			}
		}
	}

	return &recommends, comicIDs, nil
}

func (r *recommendComicRepository) GetRecommendComics1(ctx context.Context, req request.GetRecommendComicRequest, rmIds []int64) (*[]map[string]interface{}, []int64, error) {
	var recommends []map[string]interface{}
	fileds := `rc.recommend_id AS rm_id,
				cm.id AS cm_id,
				cm.name AS cm_name,
				cm.code AS cm_code,
				cm.lang AS cm_language,
				cm.audience AS cm_audience,
				cm.description AS cm_description,
				cm.cover AS cm_cover,
				cm.active AS cm_active,
				rm.position AS rm_position,
				cm.status AS cm_status,
				cm.created_at AS cm_created_at,
				cm.updated_at AS cm_updated_at,
				rm.position AS rm_position`
	tableRC := "recommend_comics as rc"

	tx := r.db.Ctx(ctx).Table(tableRC).Select(fileds).Joins("join comics as cm on rc.comic_id = cm.id ").
		Joins("join recommend_managers as rm on rc.recommend_id = rm.id ")
	if req.Language != "" {
		tx.Where("cm.lang = ?", req.Language)
	}
	tx.Where("rc.recommend_id IN (?) and cm.active = 1", rmIds)
	tx.Order("rm.created_at desc")

	err := tx.Find(&recommends)

	if err != nil {
		return nil, nil, err
	}

	var comicIDs []int64
	uniqueComicMap := make(map[int64]struct{})

	for _, recommend := range recommends {
		if cmID := util.ToInt64(recommend["cm_id"]); cmID > 0 {
			if _, exists := uniqueComicMap[cmID]; !exists {
				uniqueComicMap[cmID] = struct{}{}
				comicIDs = append(comicIDs, cmID)
			}
		}
	}

	return &recommends, comicIDs, nil
}
