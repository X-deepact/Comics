package db

import (
	"comics-admin/dto"
	"pkg-common/model"
)

type Querier interface {
	// Users
	CreateUser(user *dto.UserModel) error
	GetUserLogin(username string) (*dto.UserLoginResponse, error)
	GetUser(id int64) (*dto.UserDetailDto, error)
	GetUsers(req dto.UserListRequest) ([]*dto.UserResponse, int64, error)
	DeleteUser(id int64, adminId int64) error
	UpdateUser(user *dto.UserModel) error
	CheckUserExist(username string, phone string, email string) (*dto.UserExistDto, error)
	CheckUserExistNotMe(id int64, username string, phone string, email string) (*dto.UserExistDto, error)
	GetUserData(id int64) (*dto.UserModel, error)
	ActiveUser(id int64, adminId int64) error

	// Comics
	CreateComic(comic *model.ComicModel) error
	GetComic(id int64) (*model.ComicModel, error)
	ListComics(limit int, offset int) ([]model.ComicModel, int64, error)
	DeleteComic(id int64) error
	UpdateComic(comic *model.ComicModel) error

	// Genres
	CreateGenre(genre *model.GenreModel) error
	GetGenre(id int64) (*dto.GenreResponse, error)
	GetGenres(req dto.GenreListRequest) ([]dto.GenreResponse, int64, error)
	UpdateGenre(comic *model.GenreModel) error
	DeleteGenre(id int64) error

	// Roles
	GetRole(name string) (*model.RoleModel, error)

	// UserRoles
	CreateUserRole(userRole *model.UserRoleModel) error

	// Tiers
	GetTiers(code string) (*model.TierModel, error)
}

var _ Querier = (*Queries)(nil)
