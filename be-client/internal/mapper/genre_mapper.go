package mapper

import (
	"be-client/config"
	"be-client/internal/dto/response"
	"pkg-common/model"
)

type GenreMapper interface {
	ToModel() *model.GenreModel
	ToResponse(genres []model.GenreModel) []response.GenreResponse
	ToComicGenreResponse(genres []*model.GenreModel) []response.ComicGenreResponse
}

type genreMapper struct {
	config *config.Config
}

func NewGenreMapper(config *config.Config) GenreMapper {
	return &genreMapper{config}
}

func (m *genreMapper) ToModel() *model.GenreModel {
	return &model.GenreModel{}
}

func (m *genreMapper) ToResponse(genres []model.GenreModel) []response.GenreResponse {
	if len(genres) < 1 {
		return []response.GenreResponse{}
	}

	genreResponses := make([]response.GenreResponse, len(genres))
	for i, genre := range genres {
		genreResponses[i] = response.GenreResponse{
			Id:       genre.Id,
			Name:     genre.Name,
			Position: genre.Position,
		}
	}

	return genreResponses
}

func (m *genreMapper) ToComicGenreResponse(genres []*model.GenreModel) []response.ComicGenreResponse {
	if len(genres) == 0 {
		return []response.ComicGenreResponse{}
	}
	comicGenres := make([]response.ComicGenreResponse, len(genres))
	for i, genre := range genres {
		comicGenres[i] = response.ComicGenreResponse{
			Id:        genre.Id,
			Name:      genre.Name,
			Position:  genre.Position,
			Lang:      genre.Lang,
			CreatedBy: genre.CreatedBy,
			UpdatedBy: genre.UpdatedBy,
		}
		if genre.CreatedAt != nil {
			comicGenres[i].CreatedAt = genre.CreatedAt.UnixMilli()
		}
		if genre.UpdatedAt != nil {
			comicGenres[i].UpdatedAt = genre.UpdatedAt.UnixMilli()
		}
	}
	return comicGenres
}
