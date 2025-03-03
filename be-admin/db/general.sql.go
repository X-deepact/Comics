package db

import (
	"comics-admin/dto"
	"context"
)

func (q *Queries) GetGeneralTiers() ([]*dto.TierModel, error) {
	var tier []*dto.TierModel
	if err := q.db.WithContext(context.Background()).Table("tiers").
		Select("id, code").
		Find(&tier).Error; err != nil {
		return nil, err
	}
	return tier, nil
}

func (q *Queries) GetGeneralGenres(req dto.GeneralGenreRequest) ([]dto.GeneralGenreResponse, error) {
	var genres []dto.GeneralGenreResponse
	var query = q.db.WithContext(context.Background()).Table("genres")

	if req.Name != "" {
		query.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if req.Language != "" {
		query.Where("lang = ?", req.Language)
	}

	if err := query.Select("id, name").Order("name").Find(&genres).Error; err != nil {
		return nil, err
	}

	return genres, nil
}

func (q *Queries) GetGeneralAuthors(req dto.GeneralAuthorRequest) ([]dto.GeneralAuthorResponse, error) {
	var genres []dto.GeneralAuthorResponse
	var query = q.db.WithContext(context.Background()).Table("authors")

	if req.Name != "" {
		query.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if err := query.Select("id, name").Order("name").
		Find(&genres).Error; err != nil {
		return nil, err
	}

	return genres, nil
}
