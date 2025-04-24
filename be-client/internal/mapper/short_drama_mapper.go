package mapper

import (
	"be-client/config"
	"be-client/internal/dto/response"
	"be-client/util"
	"pkg-common/model"
)

type ShortDramaMapper interface {
	ToShortDramaByClassificationResponse(data *[]map[string]interface{}, lastEpNum *[]map[string]interface{}) []response.ShortDramaByClassificationResponse
	ToShortDramaDetailResponse(drama *map[string]interface{}, epList *[]model.EpisodeModel, genreList *[]map[string]interface{}) *response.ShortDramaDetailResponse
}
type shortDramaMapper struct {
	config *config.Config
}

func NewShortDramaMapper(config *config.Config) ShortDramaMapper {
	return &shortDramaMapper{config}
}

func (m *shortDramaMapper) ToShortDramaByClassificationResponse(data *[]map[string]interface{}, lastEpNum *[]map[string]interface{}) []response.ShortDramaByClassificationResponse {
	if len(*data) == 0 {
		return []response.ShortDramaByClassificationResponse{}
	}

	lastNumMap := make(map[int64]int64)
	for _, ep := range *lastEpNum {
		dramaID := util.ToInt64(ep["drama_id"])
		lastNum := util.ToInt64(ep["number"])
		lastNumMap[dramaID] = lastNum
	}

	shortDramaResponses := make([]response.ShortDramaByClassificationResponse, len(*data))

	for i, short_drama := range *data {
		id := util.ToInt64(short_drama["d_id"])
		if id <= 0 {
			continue
		}
		shortDramaResponses[i] = response.ShortDramaByClassificationResponse{
			Id:                  id,
			ReleaseDate:         util.ToTimePtr(short_drama["d_release_date"]),
			Thumb:               util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.DramaThumbFilePath, util.GetStringOrEmpty(short_drama["d_thumb"])),
			CreatedAt:           util.ToTimePtr(short_drama["d_created_at"]),
			UpdatedAt:           util.ToTimePtr(short_drama["d_updated_at"]),
			Name:                util.GetStringOrEmpty(short_drama["d_name"]),
			LatestEpisodeNumber: lastNumMap[id],
			Language:            util.GetStringOrEmpty(short_drama["d_tran_language"]),
			TranslatedName:      util.GetStringOrEmpty(short_drama["d_tran_title"]),
			Description:         util.GetStringOrEmpty(short_drama["d_tran_description"]),
		}
	}

	return shortDramaResponses
}

func (m *shortDramaMapper) ToShortDramaDetailResponse(drama *map[string]interface{}, epList *[]model.EpisodeModel, genreList *[]map[string]interface{}) *response.ShortDramaDetailResponse {
	if drama == nil {
		return nil
	}

	var shortDramaDetailResponse response.ShortDramaDetailResponse
	shortDramaDetailResponse.Id = util.ToInt64((*drama)["d_id"])
	shortDramaDetailResponse.Name = util.GetStringOrEmpty((*drama)["d_name"])
	shortDramaDetailResponse.TranslatedName = util.GetStringOrEmpty((*drama)["d_tran_title"])
	shortDramaDetailResponse.ReleaseDate = util.ToTimePtr((*drama)["d_release_date"])
	shortDramaDetailResponse.Thumb = util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.DramaThumbFilePath, util.GetStringOrEmpty((*drama)["d_thumb"]))
	shortDramaDetailResponse.Language = util.GetStringOrEmpty((*drama)["d_tran_language"])
	shortDramaDetailResponse.Description = util.GetStringOrEmpty((*drama)["d_tran_description"])
	shortDramaDetailResponse.CreatedAt = util.ToTimePtr((*drama)["d_created_at"])
	shortDramaDetailResponse.UpdatedAt = util.ToTimePtr((*drama)["d_updated_at"])
	shortDramaDetailResponse.Episodes = []response.EpisodeResponse{}
	shortDramaDetailResponse.Genres = []response.GenresForShortDramaResponse{}

	for _, episode := range *epList {
		episodeResponse := response.EpisodeResponse{
			Id:        episode.Id,
			Number:    episode.Number,
			CreatedAt: episode.CreatedAt,
			UpdatedAt: episode.UpdatedAt,
		}
		shortDramaDetailResponse.Episodes = append(shortDramaDetailResponse.Episodes, episodeResponse)
	}

	for _, genre := range *genreList {
		genreResponse := response.GenresForShortDramaResponse{
			Id:             util.ToInt64(genre["id"]),
			Name:           util.GetStringOrEmpty(genre["name"]),
			TranslatedName: util.GetStringOrEmpty(genre["translated_name"]),
			Position:       util.ToInt64(genre["position"]),
			Language:       util.GetStringOrEmpty(genre["language"]),
			CreatedAt:      util.ToTimePtr(genre["created_at"]),
			UpdatedAt:      util.ToTimePtr(genre["updated_at"]),
		}
		shortDramaDetailResponse.Genres = append(shortDramaDetailResponse.Genres, genreResponse)
	}

	return &shortDramaDetailResponse
}
