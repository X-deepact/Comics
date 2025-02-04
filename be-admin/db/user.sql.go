package db

import "context"

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

func (q *Queries) CreateUser(user *User) error {
	return q.db.WithContext(context.Background()).Create(user).Error
}

func (q *Queries) GetUser(username string) (*User, error) {
	var user User
	if err := q.db.WithContext(context.Background()).Where(User{Username: username}).First(&user).Error; err != nil {
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
	return q.db.WithContext(context.Background()).Where(User{Username: username}).Delete(&User{}).Error
}

func (q *Queries) UpdateUser(user *User) error {
	query := `
		UPDATE users
		SET
		    password = COALESCE(?, password),
		    email = COALESCE(?, email),
		    is_active = COALESCE(?, is_active),
		    full_name = COALESCE(?, full_name),
		    role = COALESCE(?, role)
		WHERE username = ?
	`

	return q.db.WithContext(context.Background()).Exec(query,
		nullableString(user.Password), // Handle empty or nil values
		nullableString(user.Email),
		nullableBool(user.IsActive),
		nullableString(user.FullName),
		nullableString(string(user.Role)),
		user.Username,
	).Error
}
