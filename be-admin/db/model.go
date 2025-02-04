package db

import (
	_ "gorm.io/gorm"
	"time"
)

type UsersRole string

const (
	UsersRoleUser  UsersRole = "user"
	UsersRoleAdmin UsersRole = "admin"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"type:varchar(50);not null;unique"`
	Password  string    `json:"password"`
	Email     string    `gorm:"not null;unique" json:"email"`
	FullName  string    `json:"full_name"`
	Role      UsersRole `gorm:"type:enum('user','admin')" json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
