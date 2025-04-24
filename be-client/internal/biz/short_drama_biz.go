package biz

import (
	"be-client/internal/dto/request"
	"be-client/internal/dto/response"
	"be-client/internal/mapper"
	"be-client/internal/repository"
	"context"
	"errors"
	"log/slog"
	"pkg-common/common"
	"pkg-common/model"
)

type ShortDramaBiz interface {
	GetShortDramaByClassification(ctx context.Context, req request.GetShortDramaByClassificationRequest) ([]response.ShortDramaByClassificationResponse, common.Pagination, error)
	GetEpByEpId(ctx context.Context, epId int64) (*model.EpisodeModel, error)
	GetShortDramaDetailById(ctx context.Context, id int64, req request.GetShortDramaDetailRequest) (*response.ShortDramaDetailResponse, error)
}

type shortDramaBiz struct {
	shortDramaMapper     mapper.ShortDramaMapper
	subtitleMapper       mapper.SubtitleMapper
	shortDramaRepository repository.ShortDramaRepository
	dramaTranslationRepo repository.DramaTranslationRepository
	subtitleRepo         repository.SubtitleRepository
	episodeRepo          repository.EpisodeRepository
	genreShortDramaRepo  repository.GenreForShortDramaRepository
}

func NewShortDramaBiz(
	repo repository.ShortDramaRepository,
	dramaTranslationRepo repository.DramaTranslationRepository,
	mapper mapper.ShortDramaMapper,
	episodeReposity repository.EpisodeRepository,
	genreShortDramaRepo repository.GenreForShortDramaRepository,
) ShortDramaBiz {
	return &shortDramaBiz{
		shortDramaMapper:     mapper,
		shortDramaRepository: repo,
		dramaTranslationRepo: dramaTranslationRepo,
		episodeRepo:          episodeReposity,
		genreShortDramaRepo:  genreShortDramaRepo,
	}
}

func (b *shortDramaBiz) GetShortDramaByClassification(ctx context.Context, req request.GetShortDramaByClassificationRequest) ([]response.ShortDramaByClassificationResponse, common.Pagination, error) {
	slog.Info("Get short drama by classification")
	var defaultErr = errors.New("Internal Server Error")

	results, total, dramaIds, err := b.shortDramaRepository.GetShortDramaByClassification(ctx, req)
	if err != nil {
		return []response.ShortDramaByClassificationResponse{}, common.Pagination{}, defaultErr
	}

	lastEpNum, err := b.episodeRepo.GetLatestEpNumByDranaIds(ctx, dramaIds)
	if err != nil {
		return []response.ShortDramaByClassificationResponse{}, common.Pagination{}, defaultErr
	}

	data := b.shortDramaMapper.ToShortDramaByClassificationResponse(results, lastEpNum)
	pagination := common.Pagination{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    int(total),
	}
	return data, pagination, nil
}

func (b *shortDramaBiz) GetEpByEpId(ctx context.Context, episodeId int64) (*model.EpisodeModel, error) {
	results, err := b.episodeRepo.GetById(ctx, episodeId)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (b *shortDramaBiz) GetShortDramaDetailById(ctx context.Context, id int64, req request.GetShortDramaDetailRequest) (*response.ShortDramaDetailResponse, error) {
	slog.Info("Get short drama detail by id")
	var defaultErr = errors.New("Internal Server Error")

	drama, err := b.shortDramaRepository.GetShortDramaById(ctx, id, req)
	if err != nil {
		return &response.ShortDramaDetailResponse{}, defaultErr
	}

	if len(*drama) == 0 {
		return &response.ShortDramaDetailResponse{}, errors.New("No Record Found")
	}

	epList, err := b.episodeRepo.GetAllByDramaId(ctx, id)
	if err != nil {
		return &response.ShortDramaDetailResponse{}, defaultErr
	}

	genreList, err := b.genreShortDramaRepo.GetAllGenreByDramaId(ctx, id, req)
	if err != nil {
		return &response.ShortDramaDetailResponse{}, defaultErr
	}

	return b.shortDramaMapper.ToShortDramaDetailResponse(drama, epList, genreList), nil
}
