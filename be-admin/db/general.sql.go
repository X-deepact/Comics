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

func (q *Queries) GetGeneralComics(req dto.GeneralComicRequest) ([]dto.GeneralComicResponse, error) {
	var comics []dto.GeneralComicResponse
	var query = q.db.WithContext(context.Background()).Table("comics")

	if req.Name != "" {
		query.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if req.Language != "" {
		query.Where("lang = ?", req.Language)
	}

	if req.OriginalLanguage != "" {
		query.Where("original_language = ?", req.OriginalLanguage)
	}

	if req.Audience != "" {
		query.Where("audience = ? OR audience = 'all'", req.Audience)
	}

	if req.Author != 0 {
		query = query.Joins("JOIN comic_authors ON comics.id = comic_authors.comic_id").Where("comic_authors.author_id = ?", req.Author)
	}

	if req.Genre != 0 {
		query = query.Joins("JOIN comic_genres ON comics.id = comic_genres.comic_id").Where("comic_genres.genre_id = ?", req.Genre)
	}

	if err := query.Select("comics.id, comics.name").Order("comics.name").
		Find(&comics).Error; err != nil {
		return nil, err
	}

	return comics, nil
}
