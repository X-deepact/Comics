package repository

import (
	"be-client/internal/infra/database"
	"context"
	"pkg-common/model"
)

type ProfileRepository interface {
	CreateProfile(ctx context.Context, profile *model.ProfileModel) error
	GetUserProfile(ctx context.Context, username string) (*model.ProfileModel, error)
}

type profileRepository struct {
	db database.DBInterface
}

func NewProfileRepository(db database.DBInterface) ProfileRepository {
	return &profileRepository{db: db}
}

func (r *profileRepository) CreateProfile(ctx context.Context, profile *model.ProfileModel) error {
	return r.db.Ctx(ctx).Create(profile)
}

func (r *profileRepository) GetUserProfile(ctx context.Context, username string) (*model.ProfileModel, error) {
	req := model.ProfileModel{}
	err := r.db.Ctx(ctx).
		Model(&model.ProfileModel{}).
		Joins("JOIN users u ON u.id = profiles.user_id").
		Where("u.username = ? AND profiles.active = true", username).
		Select("profiles.*").
		First(&req)
	return &req, err
}
