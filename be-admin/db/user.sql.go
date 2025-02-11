package db

import (
	"comics-admin/dto"
	"context"
	"pkg-common/model"
)

func nullableString(s string) interface{} {
	if s == "" {
		return nil
	}
	return s
}

func nullableBool(b bool) interface{} {
	// Return nil for default (zero) boolean values
	if !b {
		return nil
	}
	return b
}

func (q *Queries) CreateUser(user *model.UserModel) error {
	return q.db.WithContext(context.Background()).Create(user).Error
}

func (q *Queries) GetUser(username string) (*dto.UserModel, error) {
	var user dto.UserModel
	if err := q.db.WithContext(context.Background()).Table("users").
		Select("users.*, r.name AS role_name").
		Joins("LEFT JOIN user_roles ur ON ur.user_id = users.id").
		Joins("LEFT JOIN roles r ON r.id = ur.role_id").
		Where("users.username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	//used this one
	//When you need more control over the query.
	//For queries with dynamic conditions or involving multiple fields or expressions.
	//if err := q.db.WithContext(context.Background()).Where("username = ?", username).First(&user).Error; err != nil {
	//	return nil, err
	//}
	return &user, nil
}
func (q *Queries) DeleteUser(username string) error {
	return q.db.WithContext(context.Background()).Where(model.UserModel{Username: username}).Delete(&model.UserModel{}).Error
}

func (q *Queries) UpdateUser(user *model.UserModel) error {
	query := `
		UPDATE users
		SET
		    hash_password = COALESCE(?, hash_password),
		    email = COALESCE(?, email),
		    active = COALESCE(?, active),
		WHERE username = ?
	`

	return q.db.WithContext(context.Background()).Exec(query,
		nullableString(user.HashPassword), // Handle empty or nil values
		nullableString(user.Email),
		nullableBool(user.Active),
		user.Username,
	).Error
}

func (q *Queries) GetUserByUsername(username string) (*model.UserModel, error) {
	var user model.UserModel
	if err := q.db.WithContext(context.Background()).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
