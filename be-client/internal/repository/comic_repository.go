package repository

import (
	"be-client/internal/dto/request"
	"be-client/internal/infra/database"
	"be-client/util"
	"context"
	"fmt"
	"pkg-common/model"
)

type ComicRepository interface {
	GetPopularComicsByClassification(ctx context.Context, request request.GetPopularComicsRequest) (*[]model.ComicModel, []int64, int64, error)
	GetComicById(ctx context.Context, id int64, language string) (*model.ComicModel, error)
	GetComicsBySearch(ctx context.Context, req request.GetComicSearchRequest) (*[]model.ComicModel, []int64, int64, error)
	GetComicsByComicIds(ctx context.Context, comicIds []int64) (*[]model.ComicModel, error)
	GetComicAndDramaBySearch(ctx context.Context, req request.GetComicSearchRequest, dramaIds []int64) (*[]map[string]interface{}, []int64, int64, error)
}

type comicRepository struct {
	db database.DBInterface
}

func NewComicRepository(db database.DBInterface) ComicRepository {
	return &comicRepository{db}
}

func (r *comicRepository) GetPopularComicsByClassification(ctx context.Context, req request.GetPopularComicsRequest) (*[]model.ComicModel, []int64, int64, error) {
	var comics []model.ComicModel
	var comicIDs []int64

	baseConditions := func(query database.DBInterface) database.DBInterface {
		query = query.Joins("JOIN comic_genres ON comic_genres.comic_id = comics.id").
			Joins("JOIN genres ON genres.id = comic_genres.genre_id").
			Where("comics.active = ?", true)

		if req.Genre > 0 {
			query = query.Where("genres.id = ?", req.Genre)
		}
		if req.Progress != "" {
			query = query.Where("comics.status = ?", req.Progress)
		}
		if req.Audience != "" {
			query = query.Where("comics.audience = ?", req.Audience)
		}
		if req.Language != "" {
			query = query.Where("comics.lang = ?", req.Language)
		}

		if req.OriginalLanguage != "" {
			query = query.Where("comics.original_language = ?", req.OriginalLanguage)
		}

		return query
	}

	var total int64
	countQuery := baseConditions(r.db.Ctx(ctx).Model(&model.ComicModel{}))
	if err := countQuery.Distinct("comics.id").Count(&total); err != nil {
		return nil, comicIDs, 0, err
	}

	comicIDQuery := baseConditions(r.db.Ctx(ctx).Model(&model.ComicModel{})).
		Select("comics.*").
		Group("comics.id").
		Order("comics.updated_at DESC").
		Limit(req.PageSize).
		Offset((req.Page - 1) * req.PageSize)

	if err := comicIDQuery.Find(&comics); err != nil {
		return nil, comicIDs, 0, err
	}

	if len(comics) == 0 {
		return &comics, comicIDs, total, nil
	}

	for _, comic := range comics {
		if cmID := comic.Id; cmID > 0 {
			comicIDs = append(comicIDs, cmID)
		}
	}

	return &comics, comicIDs, total, nil
}

func (r *comicRepository) GetComicById(ctx context.Context, id int64, language string) (*model.ComicModel, error) {
	var comic model.ComicModel
	err := r.db.Ctx(ctx).
		Where("id = ?", id).
		Where("lang = ?", language).
		First(&comic)
	if err != nil {
		return nil, err
	}
	return &comic, nil
}

func (r *comicRepository) GetComicsBySearch(ctx context.Context, req request.GetComicSearchRequest) (*[]model.ComicModel, []int64, int64, error) {
	var comics []model.ComicModel
	var total int64

	limit := req.PageSize
	page := req.Page

	if page <= 0 {
		page = 1
	}

	query := r.db.Ctx(ctx).Model(&model.ComicModel{}).Where("comics.active = ?", true)
	if req.Language != "" {
		query = query.Where("comics.lang = ?", req.Language)
	}

	if req.SearchKeyword != "" {
		query = query.Where(`(comics.name LIKE ?
			OR comics.id IN
			(SELECT comic_id FROM comic_authors
			WHERE author_id IN (SELECT id FROM authors WHERE name LIKE ?) group by comic_id))`, "%"+req.SearchKeyword+"%", "%"+req.SearchKeyword+"%")
	}
	if err := query.Count(&total); err != nil {
		return nil, nil, 0, err
	}

	if limit > 0 {
		offset := (page - 1) * limit
		query = query.Limit(limit).Offset(offset)
	}
	query = query.Order("comics.updated_at DESC, comics.id ASC")
	query = query.Group("comics.id")

	err := query.Find(&comics)
	if err != nil {
		return nil, nil, 0, err
	}

	var comicIDs []int64
	for _, comic := range comics {
		if cmID := comic.Id; cmID > 0 {
			comicIDs = append(comicIDs, cmID)
		}
	}

	return &comics, comicIDs, total, nil
}

func (r *comicRepository) GetComicsByComicIds(ctx context.Context, comicIds []int64) (*[]model.ComicModel, error) {
	var comics []model.ComicModel

	query := r.db.Ctx(ctx).Model(&model.ComicModel{}).
		Where(`comics.id IN (?)`, comicIds)

	err := query.Find(&comics)
	if err != nil {
		return nil, err
	}

	return &comics, nil
}

func (r *comicRepository) GetComicAndDramaBySearch(ctx context.Context, req request.GetComicSearchRequest, dramaIds []int64) (*[]map[string]interface{}, []int64, int64, error) {
	var results []map[string]interface{}
	var total int64

	limit := req.PageSize
	page := req.Page

	if page <= 0 {
		page = 1
	}

	queryComic := r.db.Ctx(ctx).Model(&model.ComicModel{}).
		Select(`'comic' as type, comics.id, comics.name,comics.code,comics.lang,comics.audience,
		comics.description,comics.cover,comics.active,comics.created_at,comics.updated_at, null as release_date,
		null as thumb`).
		Where("comics.active = ?", true)
	if req.Language != "" {
		queryComic = queryComic.Where("comics.lang = ?", req.Language)
	}

	if req.SearchKeyword != "" {
		queryComic = queryComic.Where(`(comics.name LIKE ?
			OR comics.id IN
			(SELECT comic_id FROM comic_authors
			WHERE author_id IN (SELECT id FROM authors WHERE name LIKE ?) group by comic_id))`, "%"+req.SearchKeyword+"%", "%"+req.SearchKeyword+"%")
	}

	queryDrama := r.db.Ctx(ctx).
		Model(&model.ShortDramaModel{}).
		Select(`'drama' as type,short_dramas.id, d_tran.title AS name, null as code, d_tran.language AS lang, 
		null as audience, d_tran.description as description, null as cover, short_dramas.active, short_dramas.created_at,
		short_dramas.updated_at, short_dramas.release_date AS release_date, short_dramas.thumb as thumb`).
		Joins("JOIN drama_translations d_tran ON d_tran.drama_id = short_dramas.id").
		Where(`short_dramas.id IN (?)`, dramaIds)

	if req.Language != "" {
		queryDrama = queryDrama.Where("d_tran.language = ?", req.Language)
	} else {
		languge := util.GetDefaultLanguage()
		queryDrama = queryDrama.Where("d_tran.language = ?", languge)
	}

	sqlComic, err := queryComic.ToSQL(&model.ComicModel{})
	if err != nil {
		return nil, nil, 0, err
	}

	sqlDrama, err := queryDrama.ToSQL(&model.ShortDramaModel{})
	if err != nil {
		return nil, nil, 0, err
	}

	combinedSQL := fmt.Sprintf("(%s) UNION ALL (%s)", sqlComic, sqlDrama)

	countSQL := fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS combined_result", combinedSQL)

	if err := r.db.Ctx(ctx).Raw(countSQL).Scan(&total); err != nil {
		return nil, nil, 0, err
	}

	combinedSQL = fmt.Sprintf("%s ORDER BY updated_at DESC", combinedSQL)

	if limit > 0 {
		offset := (page - 1) * limit
		combinedSQL = fmt.Sprintf("%s LIMIT %d OFFSET %d", combinedSQL, limit, offset)
	}

	err = r.db.Ctx(ctx).Raw(combinedSQL).Scan(&results)
	if err != nil {
		return nil, nil, 0, err
	}

	var comicIDs []int64
	for _, comic := range results {
		if util.GetStringOrEmpty(comic["type"]) == "comic" {
			if cmID := util.ToInt64(comic["id"]); cmID > 0 {
				comicIDs = append(comicIDs, cmID)
			}
		}
	}

	return &results, comicIDs, total, nil
}
