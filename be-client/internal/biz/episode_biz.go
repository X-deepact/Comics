package biz

import (
	"be-client/internal/dto/response"
	"be-client/internal/mapper"
	"be-client/internal/repository"
	"context"
	"errors"
	"log/slog"
)

type EpisodeBiz interface {
	GetEpisodeById(ctx context.Context, id int64) (*response.EpisodeGetResponse, error)
}

type episodeBiz struct {
	episodeMapper     mapper.EpisodeMapper
	episodeRepository repository.EpisodeRepository
}

func NewEpisodeBiz(repo repository.EpisodeRepository, mapper mapper.EpisodeMapper) EpisodeBiz {
	return &episodeBiz{
		episodeMapper:     mapper,
		episodeRepository: repo,
	}
}

func (b *episodeBiz) GetEpisodeById(ctx context.Context, id int64) (*response.EpisodeGetResponse, error) {
	slog.Info("Get episode by Id")
	results, err := b.episodeRepository.GetById(ctx, id)
	if err != nil {
		return &response.EpisodeGetResponse{}, errors.New("Internal Server Error")
	}

	if results == nil || results.Id == 0 {
		return &response.EpisodeGetResponse{}, errors.New("No Record Found")
	}
	return b.episodeMapper.ToResponse(results), nil
}
