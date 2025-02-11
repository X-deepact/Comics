package db

import (
	"comics-admin/dto"
	"pkg-common/model"
)

type Querier interface {
	CreateUser(user *model.UserModel) error
	GetUser(username string) (*dto.UserModel, error)
	DeleteUser(username string) error
	UpdateUser(user *model.UserModel) error
	GetUserByUsername(username string) (*model.UserModel, error)

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
}

var _ Querier = (*Queries)(nil)
