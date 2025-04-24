package repository

import (
	"be-client/internal/infra/database"
	"context"
	"pkg-common/model"
)

type ChapterItemRepository interface {
	GetChapterItemsByChapterId(ctx context.Context, chapterId int64) ([]*model.ChapterItemModel, error)
}

type chapterItemRepository struct {
	db database.DBInterface
}

func NewChapterItemRepository(db database.DBInterface) ChapterItemRepository {
	return &chapterItemRepository{db: db}
}

func (r *chapterItemRepository) GetChapterItemsByChapterId(ctx context.Context, chapterId int64) ([]*model.ChapterItemModel, error) {
	var chapterItems []*model.ChapterItemModel
	err := r.db.Ctx(ctx).Where("chapter_id = ?", chapterId).
		Where("active = ?", true).
		Order("page ASC").Find(&chapterItems)
	return chapterItems, err
}
