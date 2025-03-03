package db

import (
	"comics-admin/dto"
	"context"
	"pkg-common/model"
)

// CreateAds creates new advertisement
// Hàm CreateAds dùng để tạo mới một quảng cáo
func (q *Queries) CreateAds(ads *model.AdModel) error {
	return q.db.WithContext(context.Background()).Create(ads).Error
}

// GetAds get infomation ads by id
func (q *Queries) GetAds(id int64) (*dto.AdsResponse, error) {
	var ads dto.AdsResponse
	if err := q.db.WithContext(context.Background()).
		Table("ads").
		Select("ads.*, uc.username AS created_by_name, up.username AS updated_by_name").
		Joins("LEFT JOIN users uc ON uc.id = ads.created_by").
		Joins("LEFT JOIN users up ON up.id = ads.updated_by").
		Where("ads.id = ?", id).
		First(&ads).Error; err != nil {
		return nil, err
	}
	return &ads, nil
}

// GetAdsList gets list of ads based on conditions
func (q *Queries) GetAdsList(req dto.AdsListRequest) ([]dto.AdsResponse, int64, error) {
	var ads []dto.AdsResponse
	var total int64
	query := q.db.WithContext(context.Background()).Table("ads")

	if req.Title != "" {
		query = query.Where("title LIKE ?", "%"+req.Title+"%")
	}
	if req.Type != "" {
		query = query.Where("type = ?", req.Type)
	}
	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}

	query = query.Select("ads.*, uc.username AS created_by_name, up.username AS updated_by_name").
		Joins("LEFT JOIN users uc ON uc.id = ads.created_by").
		Joins("LEFT JOIN users up ON up.id = ads.updated_by")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Limit(req.PageSize).
		Offset((req.Page - 1) * req.PageSize).
		Find(&ads).Error; err != nil {
		return nil, 0, err
	}

	return ads, total, nil
}

// UpdateAds updates advertisement information
func (q *Queries) UpdateAds(ads *model.AdModel) error {
	return q.db.WithContext(context.Background()).Model(ads).Updates(ads).Error
}

// DeleteAds deletes an advertisement
func (q *Queries) DeleteAds(id int64) error {
	return q.db.WithContext(context.Background()).Delete(&model.AdModel{Id: id}).Error
}

// UpdateAdsStatus updates the status of an advertisement
func (q *Queries) UpdateAdsStatus(id int64, status string) error {
	return q.db.WithContext(context.Background()).
		Model(&model.AdModel{}).
		Where("id = ?", id).
		Update("status", status).
		Error
}
