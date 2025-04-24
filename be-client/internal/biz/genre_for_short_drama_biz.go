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

type GenreForShortDramaBiz interface {
	GetAllGenreForShortDrama(ctx context.Context, req request.GetGenresForShortDramaRequest) ([]response.GenresForShortDramaResponse, error)
}

type genreForShortDramaBiz struct {
	genreForShortDramaMapper     mapper.GenreForShortDramaMapper
	genreForShortDramaRepository repository.GenreForShortDramaRepository
}

func NewGenreForShortDramaBiz(repo repository.GenreForShortDramaRepository, mapper mapper.GenreForShortDramaMapper) GenreForShortDramaBiz {
	return &genreForShortDramaBiz{
		genreForShortDramaMapper:     mapper,
		genreForShortDramaRepository: repo,
	}
}

func (b *genreForShortDramaBiz) GetAllGenreForShortDrama(ctx context.Context, req request.GetGenresForShortDramaRequest) ([]response.GenresForShortDramaResponse, error) {
	slog.Info("Get genres for short drama")
	results, err := b.genreForShortDramaRepository.GetAllGenreForShortDrama(ctx, req)
	if err != nil {
		return []response.GenresForShortDramaResponse{}, errors.New("Internal Server Error")
	}

	return b.genreForShortDramaMapper.ToResponse(results), nil
}
