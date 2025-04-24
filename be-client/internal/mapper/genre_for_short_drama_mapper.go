package mapper

import (
	"be-client/config"
	"be-client/internal/dto/response"
	"be-client/util"
)

type GenreForShortDramaMapper interface {
	ToResponse(data *[]map[string]interface{}) []response.GenresForShortDramaResponse
}
type genreForShortDramaMapper struct {
	config *config.Config
}

func NewGenreForShortDramaMapper(config *config.Config) GenreForShortDramaMapper {
	return &genreForShortDramaMapper{config}
}

func (m *genreForShortDramaMapper) ToResponse(data *[]map[string]interface{}) []response.GenresForShortDramaResponse {
	if len(*data) == 0 {
		return nil
	}

	genreResponses := make([]response.GenresForShortDramaResponse, len(*data))

	for i, genre := range *data {
		genreResponses[i] = response.GenresForShortDramaResponse{
			Id:             util.ToInt64(genre["id"]),
			Name:           util.GetStringOrEmpty(genre["name"]),
			TranslatedName: util.GetStringOrEmpty(genre["translated_name"]),
			Position:       util.ToInt64(genre["position"]),
			Language:       util.GetStringOrEmpty(genre["language"]),
			CreatedAt:      util.ToTimePtr(genre["created_at"]),
			UpdatedAt:      util.ToTimePtr(genre["updated_at"]),
		}
	}

	return genreResponses
}
