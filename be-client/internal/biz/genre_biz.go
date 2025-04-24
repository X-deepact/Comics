package biz

import (
	"be-client/internal/dto/response"
	"be-client/internal/mapper"
	"be-client/internal/repository"
	"errors"
	"log/slog"
)

type GenreBiz interface {
	GetAllGenres(isHomePage bool, lang string) ([]response.GenreResponse, error)
}

type genreBiz struct {
	genreMapper     mapper.GenreMapper
	genreRepository repository.GenreRepository
}

func NewGenreBiz(repo repository.GenreRepository, mapper mapper.GenreMapper) GenreBiz {
	return &genreBiz{
		genreMapper:     mapper,
		genreRepository: repo,
	}
}

func (b *genreBiz) GetAllGenres(isHomePage bool, lang string) ([]response.GenreResponse, error) {
	slog.Info("Get genres")
	genres, err := b.genreRepository.GetAllGenres(isHomePage, lang)
	if err != nil {
		return []response.GenreResponse{}, errors.New("Internal Server Error")
	}

	return b.genreMapper.ToResponse(*genres), nil
}
