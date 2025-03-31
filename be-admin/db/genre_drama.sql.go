package db

import (
	"comics-admin/dto"
	"context"
	"fmt"
	"pkg-common/model"
	"time"

	"gorm.io/gorm"
)

func (q *Queries) CreateGenreDrama(genre *model.GenreForShortDramaModel) error {
	return q.db.WithContext(context.Background()).Create(genre).Error
}

func (q *Queries) CreateGenreTranslation(translation *model.GenreTranslationModel) error {
	return q.db.WithContext(context.Background()).Create(translation).Error
}

func (q *Queries) UpdateGenreDrama(genre *model.GenreForShortDramaModel) error {
	return q.db.WithContext(context.Background()).Model(genre).Updates(map[string]interface{}{
		"name":       genre.Name,
		"position":   genre.Position,
		"updated_by": genre.UpdatedBy,
		"updated_at": genre.UpdatedAt,
	}).Error
}

func (q *Queries) DeleteGenreTranslations(genreId int64) error {
	return q.db.WithContext(context.Background()).
		Where("genre_id = ?", genreId).
		Delete(&model.GenreTranslationModel{}).Error
}

func (q *Queries) GetGenreDrama(id int64) (*dto.GenreDramaResponse, error) {
	// Lấy thông tin genre
	var genre model.GenreForShortDramaModel
	if err := q.db.WithContext(context.Background()).
		Table("genres_for_short_drama").
		Where("id = ?", id).
		First(&genre).Error; err != nil {
		return nil, err
	}

	// Lấy translations
	var translations []model.GenreTranslationModel
	if err := q.db.WithContext(context.Background()).
		Table("genre_translations").
		Where("genre_id = ?", id).
		Find(&translations).Error; err != nil {
		return nil, fmt.Errorf("error getting translations: %v", err)
	}

	// Chuyển đổi sang response format
	response := &dto.GenreDramaResponse{
		ID:        genre.Id,
		Name:      genre.Name,
		Position:  genre.Position,
		CreatedBy: genre.CreatedBy,
		UpdatedBy: genre.UpdatedBy,
	}

	if genre.CreatedAt != nil {
		response.CreatedAt = genre.CreatedAt.Format(time.RFC3339)
	}
	if genre.UpdatedAt != nil {
		response.UpdatedAt = genre.UpdatedAt.Format(time.RFC3339)
	}

	// Chuyển đổi translations
	translationResponses := make([]dto.GenreTranslationResponse, len(translations))
	for i, t := range translations {
		translationResponses[i] = dto.GenreTranslationResponse{
			ID:             t.Id,
			GenreId:        t.GenreId,
			Language:       t.Language,
			TranslatedName: t.TranslatedName,
			CreatedBy:      t.CreatedBy,
			UpdatedBy:      t.UpdatedBy,
		}
		if t.CreatedAt != nil {
			translationResponses[i].CreatedAt = t.CreatedAt.Format(time.RFC3339)
		}
		if t.UpdatedAt != nil {
			translationResponses[i].UpdatedAt = t.UpdatedAt.Format(time.RFC3339)
		}
	}

	response.Translations = translationResponses
	return response, nil
}

func (q *Queries) DeleteGenreDrama(id int64) error {
	return q.db.WithContext(context.Background()).Transaction(func(tx *gorm.DB) error {
		// Delete translations first
		if err := tx.Where("genre_id = ?", id).Delete(&model.GenreTranslationModel{}).Error; err != nil {
			return err
		}

		// Then delete the genre
		if err := tx.Delete(&model.GenreForShortDramaModel{Id: id}).Error; err != nil {
			return err
		}

		return nil
	})
}

func (q *Queries) GetGenreDramas(req dto.GenreDramaListRequest) ([]dto.GenreDramaResponse, int64, error) {
	var total int64
	var genres []model.GenreForShortDramaModel
	query := q.db.WithContext(context.Background()).Model(&model.GenreForShortDramaModel{})

	// Apply filters
	if req.Name != "" {
		query = query.Where("genres_for_short_drama.name LIKE ?", "%"+req.Name+"%")
	}

	if req.Language != "" {
		query = query.
			Joins("LEFT JOIN genre_translations ON genre_translations.genre_id = genres_for_short_drama.id").
			Where("genre_translations.language = ?", req.Language)
	}

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply sorting
	if req.SortBy != "" && req.Sort != "" {
		query = query.Order(fmt.Sprintf("genres_for_short_drama.%s %s", req.SortBy, req.Sort))
	} else {
		query = query.Order("genres_for_short_drama.position ASC")
	}

	// Apply pagination
	offset := (req.Page - 1) * req.PageSize
	if err := query.Distinct("genres_for_short_drama.*").Offset(offset).Limit(req.PageSize).Find(&genres).Error; err != nil {
		return nil, 0, err
	}

	// Convert to response format
	responses := make([]dto.GenreDramaResponse, len(genres))
	if len(genres) == 0 {
		return responses, total, nil
	}

	// Get all genre IDs
	genreIDs := make([]int64, len(genres))
	for i, genre := range genres {
		genreIDs[i] = genre.Id
		responses[i] = dto.GenreDramaResponse{
			ID:        genre.Id,
			Name:      genre.Name,
			Position:  genre.Position,
			CreatedBy: genre.CreatedBy,
			UpdatedBy: genre.UpdatedBy,
		}
		if genre.CreatedAt != nil {
			responses[i].CreatedAt = genre.CreatedAt.Format(time.RFC3339)
		}
		if genre.UpdatedAt != nil {
			responses[i].UpdatedAt = genre.UpdatedAt.Format(time.RFC3339)
		}
	}

	// Get all translations for these genres in one query
	var allTranslations []model.GenreTranslationModel
	if err := q.db.WithContext(context.Background()).
		Where("genre_id IN ?", genreIDs).
		Find(&allTranslations).Error; err != nil {
		return nil, 0, fmt.Errorf("error getting translations: %v", err)
	}

	// Create a map for quick lookup
	translationMap := make(map[int64][]dto.GenreTranslationResponse)
	for _, t := range allTranslations {
		translation := dto.GenreTranslationResponse{
			ID:             t.Id,
			GenreId:        t.GenreId,
			Language:       t.Language,
			TranslatedName: t.TranslatedName,
			CreatedBy:      t.CreatedBy,
			UpdatedBy:      t.UpdatedBy,
		}
		if t.CreatedAt != nil {
			translation.CreatedAt = t.CreatedAt.Format(time.RFC3339)
		}
		if t.UpdatedAt != nil {
			translation.UpdatedAt = t.UpdatedAt.Format(time.RFC3339)
		}
		translationMap[t.GenreId] = append(translationMap[t.GenreId], translation)
	}

	// Assign translations to respective genres
	for i, genre := range genres {
		responses[i].Translations = translationMap[genre.Id]
	}

	return responses, total, nil
}

func (q *Queries) CreateGenreTranslations(translations []*model.GenreTranslationModel) error {
	return q.db.WithContext(context.Background()).Create(&translations).Error
}
