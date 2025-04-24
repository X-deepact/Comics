package biz

import (
	"be-client/internal/dto/request"
	"be-client/internal/dto/response"
	"be-client/internal/mapper"
	"be-client/internal/repository"
	"context"
	"errors"
	"log/slog"
)

type RecommendComicBiz interface {
	GetRecommendComics(ctx context.Context, req request.GetRecommendComicRequest) ([]response.RecommendComicResponse, error)
}

type recommendComicBiz struct {
	recommendComicMapper     mapper.RecommendComicMapper
	recommendComicRepository repository.RecommendComicRepository
	recommendManagerRepo     repository.RecommendManagerRepository
	chapterRepo              repository.ChapterRepository
	chapterMapper            mapper.ChapterMapper
}

func NewRecommendComicBiz(
	repo repository.RecommendComicRepository,
	mapper mapper.RecommendComicMapper,
	recommendManagerRepo repository.RecommendManagerRepository,
	chapterRepo repository.ChapterRepository,
	chapterMapper mapper.ChapterMapper) RecommendComicBiz {
	return &recommendComicBiz{
		recommendComicRepository: repo,
		recommendComicMapper:     mapper,
		recommendManagerRepo:     recommendManagerRepo,
		chapterRepo:              chapterRepo,
		chapterMapper:            chapterMapper,
	}
}

func (u *recommendComicBiz) GetRecommendComics(ctx context.Context, req request.GetRecommendComicRequest) ([]response.RecommendComicResponse, error) {
	slog.Info("get recommend_comic")
	var defaultErr = errors.New("Internal Server Error")

	recommendList, rmIds, err := u.recommendManagerRepo.GetAllRecommend(ctx)
	if err != nil {
		return []response.RecommendComicResponse{}, defaultErr
	}

	recommendComicList, cmIds, err := u.recommendComicRepository.GetRecommendComics1(ctx, req, rmIds)

	if err != nil {
		return []response.RecommendComicResponse{}, defaultErr
	}

	comicChapterList, err := u.chapterRepo.GetLastCreatedChaptersByComicIds(ctx, cmIds)
	if err != nil {
		return []response.RecommendComicResponse{}, defaultErr
	}

	return u.recommendComicMapper.ToResponseList1(recommendList, recommendComicList, comicChapterList, u.chapterMapper), nil
}
