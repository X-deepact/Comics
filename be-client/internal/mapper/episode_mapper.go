package mapper

import (
	"be-client/config"
	"be-client/internal/dto/response"
	"be-client/util"
	"pkg-common/model"
)

type EpisodeMapper interface {
	ToResponse(data *model.EpisodeModel) *response.EpisodeGetResponse
}
type episodeMapper struct {
	config *config.Config
}

func NewEpisodeMapper(config *config.Config) EpisodeMapper {
	return &episodeMapper{config}
}

func (m *episodeMapper) ToResponse(data *model.EpisodeModel) *response.EpisodeGetResponse {
	episodeResponse := response.EpisodeGetResponse{
		Video:     util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.DramaFilePath, data.Video),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return &episodeResponse
}
