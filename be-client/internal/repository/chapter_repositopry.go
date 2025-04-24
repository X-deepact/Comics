package repository

import (
	"be-client/internal/dto/request"
	"be-client/internal/infra/database"
	"be-client/util"
	"context"
	"pkg-common/model"
	"time"

	"gorm.io/gorm"
)

type ChapterRepository interface {
	GetChaptersByComicId(ctx context.Context, comicId int64) ([]*model.ChapterModel, error)
	GetChaptersById(ctx context.Context, id int64) (*model.ChapterModel, error)
	GetNextChapterBychapterId(ctx context.Context, comicId, id int64) (int, error)
	GetPreviousChapterBychapterId(ctx context.Context, comicId, id int64) (int, error)
	GetLastCreatedChaptersWithComics(ctx context.Context, req request.GetComicRecentRequest) (*[]map[string]interface{}, []int64, int64, error)
	GetLastCreatedChaptersByComicIds(ctx context.Context, comicIds []int64) (*[]map[string]interface{}, error)
}

type chapterRepository struct {
	db database.DBInterface
}

func NewChapterRepository(db database.DBInterface) ChapterRepository {
	return &chapterRepository{db: db}
}

func (r *chapterRepository) GetChaptersByComicId(ctx context.Context, comicId int64) ([]*model.ChapterModel, error) {
	var chapters []*model.ChapterModel
	err := r.db.Ctx(ctx).Where("comic_id = ?", comicId).Find(&chapters)
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return chapters, err
}

func (r *chapterRepository) GetChaptersById(ctx context.Context, id int64) (*model.ChapterModel, error) {
	var chapter model.ChapterModel
	err := r.db.Ctx(ctx).Model(&chapter).
		Joins("JOIN comics ON comics.id = chapters.comic_id").
		Where("chapters.id = ?", id).
		Where("chapters.active = ?", true).
		Select("chapters.*").
		Limit(1).
		Find(&chapter)
	return &chapter, err
}

func (r *chapterRepository) GetNextChapterBychapterId(ctx context.Context, comicId, id int64) (int, error) {
	var chapter model.ChapterModel
	err := r.db.Ctx(ctx).Model(&chapter).
		Where("comic_id = ?", comicId).
		Where("id > ?", id).
		Order("number ASC").
		Limit(1).
		First(&chapter)
	if err != nil {
		return 0, err
	}
	return int(chapter.Id), nil
}

func (r *chapterRepository) GetPreviousChapterBychapterId(ctx context.Context, comicId, id int64) (int, error) {
	var chapter model.ChapterModel
	err := r.db.Ctx(ctx).Model(&chapter).
		Where("comic_id = ?", comicId).
		Where("id < ?", id).
		Order("number DESC").
		Limit(1).
		First(&chapter)
	if err != nil {
		return 0, err
	}
	return int(chapter.Id), nil
}

func (r *chapterRepository) GetLastCreatedChaptersWithComics(ctx context.Context, req request.GetComicRecentRequest) (*[]map[string]interface{}, []int64, int64, error) {
	var chapters []map[string]interface{}
	var total int64

	limit := req.PageSize
	page := req.Page

	if page <= 0 {
		page = 1
	}

	query := r.db.Ctx(ctx).Model(&model.ChapterModel{}).
		Select(`chapters.comic_id as cm_id, MAX(chapters.created_at) AS last_created_date`).
		Joins(`Join comics cm ON cm.id = chapters.comic_id`).
		Where(`chapters.active_from < ? AND chapters.active = ? `, time.Now(), true).
		Where("cm.active = ?", true).
		Group(`chapters.comic_id`).
		Order(`MAX(chapters.created_at) desc`)

	if req.Language != "" {
		query = query.Where("cm.lang = ?", req.Language)
	}

	if err := query.Count(&total); err != nil {
		return nil, nil, 0, err
	}

	if limit > 0 {
		offset := (page - 1) * limit
		query = query.Limit(limit).Offset(offset)
	}

	err := query.Find(&chapters)
	if err != nil {
		return nil, nil, 0, err
	}

	var comicIDs []int64
	idSet := make(map[int64]struct{})

	for _, chapter := range chapters {
		if cmID := util.ToInt64(chapter["cm_id"]); cmID > 0 {
			if _, exists := idSet[cmID]; !exists {
				idSet[cmID] = struct{}{}
				comicIDs = append(comicIDs, cmID)
			}
		}
	}
	return &chapters, comicIDs, total, err
}

func (r *chapterRepository) GetLastCreatedChaptersByComicIds(ctx context.Context, comicIds []int64) (*[]map[string]interface{}, error) {
	var chapters []map[string]interface{}
	limit := 3

	query := `SELECT * FROM (
				SELECT comic_id as cm_id, created_at as ch_created_at, name as ch_name, id as ch_id, cover as ch_cover, number as ch_number,
					   ROW_NUMBER() OVER (PARTITION BY comic_id ORDER BY created_at DESC, id DESC) AS rn
				FROM chapters
				WHERE active_from < ? AND active = ? AND comic_id IN (?)
			) ch_filtered
			WHERE rn <= ?`
	err := r.db.Ctx(ctx).
		Raw(query, time.Now(), true, comicIds, limit).Find(&chapters)

	if err != nil {
		return nil, err
	}
	return &chapters, err
}
