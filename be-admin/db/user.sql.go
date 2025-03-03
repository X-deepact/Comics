package db

import (
	"comics-admin/dto"
	"context"
	"fmt"
	"pkg-common/model"
	"time"

	"gorm.io/gorm"
)

func (q *Queries) CreateUser(user *dto.UserModel) error {
	return q.db.WithContext(context.Background()).Create(user).Error
}

func (q *Queries) GetUserLogin(username string) (*dto.UserLoginResponse, error) {
	var user dto.UserLoginResponse
	if err := q.db.WithContext(context.Background()).Table("users u").
		Select("u.id, u.username, u.phone, u.email, u.birth_day AS birthday, u.first_name, u.last_name, u.full_name, u.active, u.hash_password, r.name AS role_name,"+
			"p.display_name, p.description, p.avatar, p.tier_id, t.code AS tier_code, u.created_at, uc.username AS created_by_name, u.updated_at, up.username AS updated_by_name").
		Joins("LEFT JOIN user_roles ur ON ur.user_id = u.id").
		Joins("LEFT JOIN roles r ON r.id = ur.role_id").
		Joins("LEFT JOIN profiles p ON p.user_id = u.id").
		Joins("LEFT JOIN tiers t ON t.id = p.tier_id").
		Joins("LEFT JOIN users uc ON uc.id = u.created_by").
		Joins("LEFT JOIN users up ON up.id = u.updated_by").
		Where("(u.username = @username or u.email = @username) and u.active = true and u.deleted_at is null",
			map[string]interface{}{
				"username": username,
			}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (q *Queries) GetUser(id int64) (*dto.UserDetailDto, error) {
	var user dto.UserDetailDto
	if err := q.db.WithContext(context.Background()).Table("users u").
		Select("u.id, u.username, u.phone, u.email, u.birth_day AS birthday, u.first_name, u.last_name, u.full_name, u.active, r.name AS role_name,"+
			"p.display_name, p.description, p.avatar, p.tier_id, t.code AS tier_code").
		Joins("LEFT JOIN user_roles ur ON ur.user_id = u.id").
		Joins("LEFT JOIN roles r ON r.id = ur.role_id").
		Joins("LEFT JOIN profiles p ON p.user_id = u.id").
		Joins("LEFT JOIN tiers t ON t.id = p.tier_id").
		Where("u.id = ? and u.deleted_at is null", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (q *Queries) GetUsers(req dto.UserListRequest) ([]*dto.UserResponse, int64, error) {
	var users []*dto.UserResponse
	var total int64
	var query = q.db.WithContext(context.Background()).Table("users u")

	if req.Username != "" {
		query.Where("u.username LIKE ?", "%"+req.Username+"%")
	}

	if req.Phone != "" {
		query.Where("u.phone LIKE ?", "%"+req.Phone+"%")
	}

	if req.Email != "" {
		query.Where("u.email LIKE ?", "%"+req.Email+"%")
	}

	if req.Name != "" {
		query.Where("u.full_name LIKE ?", "%"+req.Name+"%")
	}

	if req.DisplayName != "" {
		query.Where("p.display_name LIKE ?", "%"+req.DisplayName+"%")
	}

	if req.TierId > 0 {
		query.Where("p.tier_id = ?", req.TierId)
	}

	if req.SortBy != "" {
		order := fmt.Sprintf("%s %s", req.SortBy, req.Sort)
		query = query.Order(order)
	}

	query = query.Select("u.id, u.username, u.phone, u.email, u.birth_day AS birthday, u.first_name, u.last_name, u.full_name, u.active, r.name AS role_name," +
		"p.display_name, p.description, p.avatar, p.tier_id, t.code AS tier_code, u.created_at, uc.username AS created_by_name, u.updated_at, up.username AS updated_by_name").
		Joins("LEFT JOIN user_roles ur ON ur.user_id = u.id").
		Joins("LEFT JOIN roles r ON r.id = ur.role_id").
		Joins("LEFT JOIN profiles p ON p.user_id = u.id").
		Joins("LEFT JOIN tiers t ON t.id = p.tier_id").
		Joins("LEFT JOIN users uc ON uc.id = u.created_by").
		Joins("LEFT JOIN users up ON up.id = u.updated_by").
		Where("u.deleted_at is null")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("id DESC").
		Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).
		Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (q *Queries) DeleteUser(id int64, adminId int64) error {
	deletedUser := map[string]interface{}{
		"deleted_at": time.Now(),
		"deleted_by": adminId,
	}

	return q.db.WithContext(context.Background()).Model(&model.UserModel{}).Where("id = ?", id).Updates(deletedUser).Error
}

func (q *Queries) GetUserData(id int64) (*dto.UserModel, error) {
	var user dto.UserModel
	if err := q.db.WithContext(context.Background()).First(&user, id).Error; err != nil {
		return nil, err
	}

	var profile model.ProfileModel
	if err := q.db.WithContext(context.Background()).Where("user_id = ?", id).First(&profile).Error; err != nil {
		return nil, err
	}

	user.Profile = profile

	return &user, nil
}

func (q *Queries) UpdateUser(user *dto.UserModel) error {
	// Start transaction
	tx := q.db.WithContext(context.Background()).Begin()

	// Update user table
	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update profile separately if it exists
	if user.Profile.Id != 0 {
		if err := tx.Save(user.Profile).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	return tx.Commit().Error
}

func (q *Queries) CheckUserExist(username string, phone *string, email *string) (*dto.UserExistDto, error) {
	var userExist dto.UserExistDto

	query := `
		SELECT MAX(username = @username) AS is_username, MAX(phone = @phone) AS is_phone, MAX(email = @email) AS is_email
		FROM users
		WHERE (username = @username or if(@phone is null, false, phone = @phone) or if(@email is null, false, email = @email)) and deleted_at is null
	`
	if err := q.db.WithContext(context.Background()).Raw(query,
		map[string]interface{}{
			"username": username,
			"phone":    phone,
			"email":    email,
		}).Scan(&userExist).Error; err != nil {
		return nil, err
	}
	return &userExist, nil
}

func (q *Queries) CheckUserExistNotMe(id int64, username string, phone *string, email *string) (*dto.UserExistDto, error) {
	var userExist dto.UserExistDto

	query := `
		SELECT MAX(username = @username) AS is_username, MAX(phone = @phone) AS is_phone, MAX(email = @email) AS is_email
		FROM users
		WHERE id != @id and (username = @username or if(@phone is null, false, phone = @phone) or if(@email is null, false, email = @email)) and deleted_at is null
	`
	if err := q.db.WithContext(context.Background()).Raw(query,
		map[string]interface{}{
			"id":       id,
			"username": username,
			"phone":    phone,
			"email":    email,
		}).Scan(&userExist).Error; err != nil {
		return nil, err
	}
	return &userExist, nil
}

func (q *Queries) ActiveUser(id int64, adminId int64) error {
	activeUser := map[string]interface{}{
		"active":     gorm.Expr("NOT active"),
		"updated_by": adminId,
	}

	return q.db.WithContext(context.Background()).Model(&model.UserModel{}).Where("id = ?", id).Updates(activeUser).Error
}

func (q *Queries) ChangePassword(id int64, password string) error {
	changePassword := map[string]interface{}{
		"hash_password": password,
		"updated_by":    id,
	}

	return q.db.WithContext(context.Background()).Model(&model.UserModel{}).Where("id = ?", id).Updates(changePassword).Error
}

func (q *Queries) GetUserNamesByIds(ids []int64) (map[int64]string, error) {
	var users []model.UserModel
	if err := q.db.WithContext(context.Background()).Where("id in (?)", ids).Find(&users).Error; err != nil {
		return nil, err
	}
	names := make(map[int64]string)
	for _, user := range users {
		names[user.Id] = user.Username
	}
	return names, nil
}
