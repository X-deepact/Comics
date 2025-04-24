package biz

import (
	"be-client/internal/dto/response"
	"be-client/internal/mapper"
	"be-client/internal/repository"
	"context"
	"errors"

	"gorm.io/gorm"
)

type ChapterBiz interface {
	GetChapterDetailByChapterId(ctx context.Context, chapterId int64, lang string) (*response.ChapterDetailResponse, error)
}

type chapterBiz struct {
	chapterRepository     repository.ChapterRepository
	chapterItemRepository repository.ChapterItemRepository
	chapterMapper         mapper.ChapterMapper
	comicRepository       repository.ComicRepository
}

func NewChapterBiz(chapterRepo repository.ChapterRepository, chapterItemRepo repository.ChapterItemRepository, chapterMapper mapper.ChapterMapper, comicRepo repository.ComicRepository) ChapterBiz {
	return &chapterBiz{chapterRepository: chapterRepo, chapterItemRepository: chapterItemRepo, chapterMapper: chapterMapper, comicRepository: comicRepo}
}

func (b *chapterBiz) GetChapterDetailByChapterId(ctx context.Context, chapterId int64, lang string) (*response.ChapterDetailResponse, error) {
	chapter, err := b.chapterRepository.GetChaptersById(ctx, chapterId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if chapter == nil {
		return nil, errors.New("chapter not found")
	}

	chapterDetail := b.chapterMapper.ToChapterDetailResponse(chapter)

	nextChapter, err := b.chapterRepository.GetNextChapterBychapterId(ctx, chapter.ComicId, chapterId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if nextChapter > 0 {
		chapterDetail.NextChapter = &nextChapter
	}
	previousChapter, err := b.chapterRepository.GetPreviousChapterBychapterId(ctx, chapter.ComicId, chapterId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if previousChapter > 0 {
		chapterDetail.PrevChapter = &previousChapter
	}

	comic, err := b.comicRepository.GetComicById(ctx, chapter.ComicId, lang)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if comic != nil {
		chapterDetail.ComicName = comic.Name
	}

	chapterItems, err := b.chapterItemRepository.GetChapterItemsByChapterId(ctx, chapterId)
	if err != nil {
		return nil, err
	}

	if len(chapterItems) > 0 {
		chapterDetail.ChapterItems = make([]response.ChapterItemResponse, len(chapterItems))
		for i := range chapterItems {
			chapterDetail.ChapterItems[i] = b.chapterMapper.ToChapterItemResponse(chapterItems[i])
		}
	}
	return &chapterDetail, nil
}
