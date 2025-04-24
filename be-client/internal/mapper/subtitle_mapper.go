package mapper

import (
	"be-client/config"
	"be-client/internal/dto/response"
	"be-client/util"
	"pkg-common/model"
)

type SubtitleMapper interface {
	MapToResponse(subtitles []model.SubtitleModel) []response.SubtitleResponse
}

type subtitleMapper struct {
	config *config.Config
}

func NewSubtitleMapper(config *config.Config) SubtitleMapper {
	return &subtitleMapper{config}
}

func (m *subtitleMapper) MapToResponse(subtitles []model.SubtitleModel) []response.SubtitleResponse {
	mapped := make([]response.SubtitleResponse, len(subtitles))
	for i, subtitle := range subtitles {
		mapped[i] = response.SubtitleResponse{
			EpisodeId:   subtitle.EpisodeId,
			Language:    subtitle.Language,
			SubtitleUrl: util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.DramaSubFilePath, subtitle.SubtitleUrl),
		}
	}
	return mapped
}
