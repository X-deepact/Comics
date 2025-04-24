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

type AdsBiz interface {
	GetAdsComics(ctx context.Context, req request.GetAdsComicRequest) ([]response.AdsComicResponse, error)
}

type adsBiz struct {
	adsMapper     mapper.AdsMapper
	adsRepository repository.AdsRepository
}

func NewAdsBiz(repo repository.AdsRepository, mapper mapper.AdsMapper) AdsBiz {
	return &adsBiz{
		adsMapper:     mapper,
		adsRepository: repo,
	}
}

func (b *adsBiz) GetAdsComics(ctx context.Context, req request.GetAdsComicRequest) ([]response.AdsComicResponse, error) {
	slog.Info("Get ads comics")
	results, err := b.adsRepository.GetAdsComics(ctx, req)
	if err != nil {
		return []response.AdsComicResponse{}, errors.New("Internal Server Error")
	}
	return b.adsMapper.ToAdsComicResponse(results), nil
}
