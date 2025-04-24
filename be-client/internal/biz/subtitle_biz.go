package biz

import (
	"be-client/internal/dto/response"
	"be-client/internal/mapper"
	"be-client/internal/repository"
	"context"
)

type SubtitleBiz interface {
	GetDramaVttFiles(ctx context.Context, episodeId int64) ([]response.SubtitleResponse, error)
}

type subtitleBiz struct {
	subtitleMapper mapper.SubtitleMapper
	subtitleRepo   repository.SubtitleRepository
}

func NewSubtitleBiz(
	subtitleMapper mapper.SubtitleMapper,
	subtitleRepo repository.SubtitleRepository,
) SubtitleBiz {
	return &subtitleBiz{
		subtitleMapper: subtitleMapper,
		subtitleRepo:   subtitleRepo,
	}
}

func (b *subtitleBiz) GetDramaVttFiles(ctx context.Context, episodeId int64) ([]response.SubtitleResponse, error) {
	//TEMPORARY, MAYBE WILL FIND EPISODE ID BY DRAMAID AND EPISODE NUMBER
	files, err := b.subtitleRepo.GetVTTsByEpisodeId(ctx, episodeId)
	if err != nil {
		return nil, err
	}
	return b.subtitleMapper.MapToResponse(files), nil
}
