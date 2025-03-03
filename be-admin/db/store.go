package db

import (
	"comics-admin/dto"
	"pkg-common/model"

	"gorm.io/gorm"
)

type Store interface {
	Querier
	CreateAds(ads *model.AdModel) error
	GetAds(id int64) (*dto.AdsResponse, error)
	GetAdsList(req dto.AdsListRequest) ([]dto.AdsResponse, int64, error)
	UpdateAds(ads *model.AdModel) error
	DeleteAds(id int64) error
	UpdateAdsStatus(id int64, status string) error
}

type SQLStore struct {
	*Queries
	DB *gorm.DB
}

func NewStore(db *gorm.DB) *SQLStore {
	return &SQLStore{
		Queries: New(db),
		DB:      db,
	}
}
