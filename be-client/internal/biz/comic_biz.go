package biz

import (
	"be-client/internal/dto/request"
	"be-client/internal/dto/response"
	"be-client/internal/mapper"
	"be-client/internal/repository"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"pkg-common/common"

	"gorm.io/gorm"
)

type ComicBiz interface {
	GetPopularComicsByClassification(ctx context.Context, req request.GetPopularComicsRequest) ([]response.ComicWithChapterResponse, common.Pagination, error)
	GetComicDetailById(context.Context, int64, string) (response.ComicDetailResponse, error)
	GetComicsBySearch(ctx context.Context, req request.GetComicSearchRequest) ([]response.ComicWithChapterAndAuthorResponse, common.Pagination, error)
	GetComicsByRecently(ctx context.Context, req request.GetComicRecentRequest) ([]response.ComicRecent, common.Pagination, error)
	GetComicAndDramaBySearch(ctx context.Context, req request.GetComicSearchRequest) ([]response.ComicAndDramaSearchResponse, common.Pagination, error)
}

type comicBiz struct {
	comicMapper     mapper.ComicMapper
	comicRepository repository.ComicRepository
	chapterRepo     repository.ChapterRepository
	chapterMapper   mapper.ChapterMapper

	authorMapper     mapper.AuthorMapper
	authorRepository repository.AuthorRepository
	comicAuthorRepo  repository.ComicAuthorRepository

	genreMapper     mapper.GenreMapper
	genreRepository repository.GenreRepository

	shortDramaRepository repository.ShortDramaRepository
	dramaTranslationRepo repository.DramaTranslationRepository
	episodeRepo          repository.EpisodeRepository
}

func NewComicBiz(repo repository.ComicRepository,
	mapper mapper.ComicMapper,
	chapterRepo repository.ChapterRepository,
	chapterMapper mapper.ChapterMapper,
	authorRepo repository.AuthorRepository,
	comicAuthorRepo repository.ComicAuthorRepository,
	authorMapper mapper.AuthorMapper,
	genreRepo repository.GenreRepository,
	genreMapper mapper.GenreMapper,
	shortDramaRepository repository.ShortDramaRepository,
	dramaTranslationRepo repository.DramaTranslationRepository,
	episodeRepo repository.EpisodeRepository) ComicBiz {
	return &comicBiz{
		comicMapper:          mapper,
		authorMapper:         authorMapper,
		authorRepository:     authorRepo,
		comicRepository:      repo,
		chapterRepo:          chapterRepo,
		comicAuthorRepo:      comicAuthorRepo,
		chapterMapper:        chapterMapper,
		genreRepository:      genreRepo,
		genreMapper:          genreMapper,
		shortDramaRepository: shortDramaRepository,
		dramaTranslationRepo: dramaTranslationRepo,
		episodeRepo:          episodeRepo,
	}
}

func (b *comicBiz) GetPopularComicsByClassification(ctx context.Context, req request.GetPopularComicsRequest) ([]response.ComicWithChapterResponse, common.Pagination, error) {
	slog.Info("Get comics")
	var defaultErr = errors.New("Internal Server Error")

	comics, comicIds, total, err := b.comicRepository.GetPopularComicsByClassification(ctx, req)
	if err != nil {
		return []response.ComicWithChapterResponse{}, common.Pagination{}, defaultErr
	}

	comicChapterList, err := b.chapterRepo.GetLastCreatedChaptersByComicIds(ctx, comicIds)
	if err != nil {
		return []response.ComicWithChapterResponse{}, common.Pagination{}, defaultErr
	}

	data := b.comicMapper.ToResponse(b.chapterMapper, b.authorMapper, comics, comicChapterList)
	pagination := common.Pagination{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    int(total),
	}
	return data, pagination, nil
}

func (b *comicBiz) GetComicDetailById(ctx context.Context, id int64, language string) (response.ComicDetailResponse, error) {
	slog.Info("Get comic detail by code")
	comic, err := b.comicRepository.GetComicById(ctx, id, language)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get comic detail by code, %v", err))
		return response.ComicDetailResponse{}, errors.New("failed to get comic detail by code")
	}
	if comic == nil {
		return response.ComicDetailResponse{}, gorm.ErrRecordNotFound
	}
	resp := b.comicMapper.ToComicDetailResponse(comic)

	authors, err := b.authorRepository.GetAuthorsByComicIds(ctx, []int64{comic.Id})
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get comic detail by code, %v", err))
		return resp, errors.New("failed to get comic detail by code")
	}

	if len(authors) > 0 {
		resp.Authors = b.authorMapper.ToAuthorComicResponse(authors)
	}

	chapters, err := b.chapterRepo.GetChaptersByComicId(ctx, comic.Id)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get comic detail by code, %v", err))
		return resp, errors.New("failed to get comic detail by code")
	}

	if len(chapters) >= 0 {
		chapterComics := make([]response.ComicChapterResponse, len(chapters))
		for i, chapter := range chapters {
			chapterComics[i] = b.chapterMapper.ToResponse(chapter)
		}
		resp.Chapters = chapterComics
	}

	genres, err := b.genreRepository.GetGenresByComicId(ctx, comic.Id)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get comic detail by code, %v", err))
		return resp, err
	}

	if len(genres) > 0 {
		resp.Genres = b.genreMapper.ToComicGenreResponse(genres)
	}

	return resp, nil
}

func (b *comicBiz) GetComicsBySearch(ctx context.Context, req request.GetComicSearchRequest) ([]response.ComicWithChapterAndAuthorResponse, common.Pagination, error) {
	slog.Info("Get comics search")
	var defaultErr = errors.New("Internal Server Error")

	results, uniqueComicIds, total, err := b.comicRepository.GetComicsBySearch(ctx, req)
	if err != nil {
		return []response.ComicWithChapterAndAuthorResponse{}, common.Pagination{}, defaultErr
	}

	comicChapterList, err := b.chapterRepo.GetLastCreatedChaptersByComicIds(ctx, uniqueComicIds)
	if err != nil {
		return []response.ComicWithChapterAndAuthorResponse{}, common.Pagination{}, defaultErr
	}

	comicAuthorResults, err := b.comicAuthorRepo.GetComicAndAuthorsByComicId(ctx, uniqueComicIds)
	if err != nil {
		return []response.ComicWithChapterAndAuthorResponse{}, common.Pagination{}, defaultErr
	}

	data := b.comicMapper.ToComicSearchResponse(b.chapterMapper, b.authorMapper, results, comicAuthorResults, comicChapterList)
	pagination := common.Pagination{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    int(total),
	}
	return data, pagination, nil
}

func (b *comicBiz) GetComicsByRecently(ctx context.Context, req request.GetComicRecentRequest) ([]response.ComicRecent, common.Pagination, error) {
	slog.Info("Get recently comics")
	var defaultErr = errors.New("Internal Server Error")

	lastChapterInfo, comicIds, total, err := b.chapterRepo.GetLastCreatedChaptersWithComics(ctx, req)
	if err != nil {
		return []response.ComicRecent{}, common.Pagination{}, defaultErr
	}

	comicChapterList, err := b.chapterRepo.GetLastCreatedChaptersByComicIds(ctx, comicIds)
	if err != nil {
		return []response.ComicRecent{}, common.Pagination{}, defaultErr
	}

	results, err := b.comicRepository.GetComicsByComicIds(ctx, comicIds)
	if err != nil {
		return []response.ComicRecent{}, common.Pagination{}, defaultErr
	}

	data := b.comicMapper.ToComicRecentResponse(b.chapterMapper, results, lastChapterInfo, comicChapterList)
	pagination := common.Pagination{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    int(total),
	}
	return data, pagination, nil
}

func (b *comicBiz) GetComicAndDramaBySearch(ctx context.Context, req request.GetComicSearchRequest) ([]response.ComicAndDramaSearchResponse, common.Pagination, error) {
	slog.Info("Get comics and drama search")
	var defaultErr = errors.New("Internal Server Error")

	dramaIds, err := b.dramaTranslationRepo.SearchDramaFromName(ctx, req)
	if err != nil {
		return []response.ComicAndDramaSearchResponse{}, common.Pagination{}, defaultErr
	}

	results, uniqueComicIds, total, err := b.comicRepository.GetComicAndDramaBySearch(ctx, req, dramaIds)
	if err != nil {
		return []response.ComicAndDramaSearchResponse{}, common.Pagination{}, defaultErr
	}

	lastEpNum, err := b.episodeRepo.GetLatestEpNumByDranaIds(ctx, dramaIds)
	if err != nil {
		return []response.ComicAndDramaSearchResponse{}, common.Pagination{}, defaultErr
	}

	comicChapterList, err := b.chapterRepo.GetLastCreatedChaptersByComicIds(ctx, uniqueComicIds)
	if err != nil {
		return []response.ComicAndDramaSearchResponse{}, common.Pagination{}, defaultErr
	}

	comicAuthorResults, err := b.comicAuthorRepo.GetComicAndAuthorsByComicId(ctx, uniqueComicIds)
	if err != nil {
		return []response.ComicAndDramaSearchResponse{}, common.Pagination{}, defaultErr
	}

	data := b.comicMapper.ToComicAndDramaSearchResponse(b.chapterMapper, b.authorMapper, results, comicAuthorResults, comicChapterList, lastEpNum)
	pagination := common.Pagination{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    int(total),
	}
	return data, pagination, nil
}
