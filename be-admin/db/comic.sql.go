package db

import (
	"comics-admin/dto"
	"context"
	"fmt"
	"pkg-common/common"
	"pkg-common/model"

	"gorm.io/gorm"
)

func (q *Queries) CreateComic(req *dto.ComicRequest) (*model.ComicModel, error) {

	comic := &model.ComicModel{
		Name:             req.Name,
		Code:             req.Code,
		Lang:             req.Lang,
		OriginalLanguage: req.OriginalLanguage,
		Audience:         req.Audience,
		Description:      req.Description,
		Cover:            req.Cover,
		CreatedBy:        req.CreatedBy,
	}

	if req.Status == common.COMIC_STATUS_COMPLETED || req.Status == common.COMIC_STATUS_ONGOING {
		comic.Status = req.Status
	}

	if err := q.db.WithContext(context.Background()).Create(comic).Error; err != nil {
		return nil, err
	}

	var comicAuthors []model.ComicAuthorModel
	for _, authorID := range req.Authors {
		comicAuthors = append(comicAuthors, model.ComicAuthorModel{
			ComicId:  comic.Id,
			AuthorId: authorID,
		})
	}

	var comicGenres []model.ComicGenreModel
	for _, genreID := range req.Genres {
		comicGenres = append(comicGenres, model.ComicGenreModel{
			ComicId: comic.Id,
			GenreId: genreID,
		})
	}

	if len(comicAuthors) > 0 {
		if err := q.db.WithContext(context.Background()).Create(&comicAuthors).Error; err != nil {
			return nil, err
		}
	}

	if len(comicGenres) > 0 {
		if err := q.db.WithContext(context.Background()).Create(&comicGenres).Error; err != nil {
			return nil, err
		}
	}

	return comic, nil
}

func (q *Queries) DeleteComic(id int64) error {
	return q.db.WithContext(context.Background()).Delete(&model.ComicModel{Id: id}).Error
}

func (q *Queries) GetComic(id int64) (*dto.ComicResponse, error) {
	var comic dto.ComicResponse

	query := q.db.WithContext(context.Background()).Table("comics")

	if err := query.Where("comics.id = ?", id).First(&comic).Error; err != nil {
		return nil, err
	}

	return &comic, nil
}

func (q *Queries) ListComics(req dto.ComicListRequest) ([]dto.ComicResponse, int64, error) {
	var comics []dto.ComicResponse
	var total int64

	query := q.db.WithContext(context.Background()).Table("comics")
	if req.Query != "" {
		query = query.Where("comics.name LIKE ? OR comics.code LIKE ?", "%"+req.Query+"%", "%"+req.Query+"%")
	}

	if req.ActiveValue {
		query = query.Where("comics.active = ?", req.Active)
	}

	if req.Lang != "" {
		query = query.Where("comics.lang = ?", req.Lang)
	}
	if req.Audience != "" {
		query = query.Where("comics.audience = ?", req.Audience)
	}

	if req.Author != 0 {
		query = query.Joins("JOIN comic_authors ON comics.id = comic_authors.comic_id").Where("comic_authors.author_id = ?", req.Author)
	}

	if req.Genre != 0 {
		query = query.Joins("JOIN comic_genres ON comics.id = comic_genres.comic_id").Where("comic_genres.genre_id = ?", req.Genre)
	}

	if req.SortBy != "" {
		order := fmt.Sprintf("%s %s", req.SortBy, req.Sort)
		query = query.Order(order)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.PageSize
	query = query.Offset(offset).Limit(req.PageSize)

	if err := query.Find(&comics).Error; err != nil {
		return nil, 0, err
	}

	return comics, total, nil
}

func (q *Queries) UpdateComic(req *dto.ComicUpdateRequest) (*dto.ComicResponse, error) {
	comic := &model.ComicModel{
		Id:               req.ID,
		Name:             req.Name,
		Code:             req.Code,
		Lang:             req.Lang,
		OriginalLanguage: req.OriginalLanguage,
		Audience:         req.Audience,
		Description:      req.Description,
		Cover:            req.Cover,
		Status:           req.Status,
		UpdatedBy:        req.UpdatedBy,
	}

	if err := q.db.WithContext(context.Background()).Model(comic).Updates(comic).Error; err != nil {
		return nil, err
	}

	if err := q.db.WithContext(context.Background()).Where("comic_id = ?", req.ID).Delete(&model.ComicAuthorModel{}).Error; err != nil {
		return nil, err
	}
	var comicAuthors []model.ComicAuthorModel
	for _, authorID := range req.Authors {
		comicAuthors = append(comicAuthors, model.ComicAuthorModel{
			ComicId:  req.ID,
			AuthorId: authorID,
		})
	}
	if len(comicAuthors) > 0 {
		if err := q.db.WithContext(context.Background()).Create(&comicAuthors).Error; err != nil {
			return nil, err
		}
	}

	if err := q.db.WithContext(context.Background()).Where("comic_id = ?", req.ID).Delete(&model.ComicGenreModel{}).Error; err != nil {
		return nil, err
	}
	var comicGenres []model.ComicGenreModel
	for _, genreID := range req.Genres {
		comicGenres = append(comicGenres, model.ComicGenreModel{
			ComicId: req.ID,
			GenreId: genreID,
		})
	}
	if len(comicGenres) > 0 {
		if err := q.db.WithContext(context.Background()).Create(&comicGenres).Error; err != nil {
			return nil, err
		}
	}

	return q.GetComic(req.ID)
}

func (q *Queries) ActiveComic(id int64, adminId int64) error {
	activeComic := map[string]interface{}{
		"active":     gorm.Expr("NOT active"),
		"updated_by": adminId,
	}

	return q.db.WithContext(context.Background()).Model(&model.ComicModel{}).Where("id = ?", id).Updates(activeComic).Error
}
