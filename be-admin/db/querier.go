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
	CheckUserExist(username string, phone *string, email *string) (*dto.UserExistDto, error)
	CheckUserExistNotMe(id int64, username string, phone *string, email *string) (*dto.UserExistDto, error)
	GetUserData(id int64) (*dto.UserModel, error)
	ActiveUser(id int64, adminId int64) error
	ChangePassword(id int64, password string) error
	GetUserNamesByIds(ids []int64) (map[int64]string, error)

	// Comics
	CreateComic(comic *dto.ComicRequest) (*model.ComicModel, error)
	GetComic(id int64) (*dto.ComicResponse, error)
	ListComics(req dto.ComicListRequest) ([]dto.ComicResponse, int64, error)
	DeleteComic(id int64) error
	UpdateComic(comic *dto.ComicUpdateRequest) (*dto.ComicResponse, error)
	ActiveComic(id int64, adminId int64) error

	// Genres
	CreateGenre(genre *model.GenreModel) error
	GetGenre(id int64) (*dto.GenreResponse, error)
	GetGenres(req dto.GenreListRequest) ([]dto.GenreResponse, int64, error)
	UpdateGenre(comic *model.GenreModel) error
	DeleteGenre(id int64) error
	GetGenresOfAComic(comicId int64) ([]dto.GenreResponse, error)

	// Roles
	GetRole(name string) (*model.RoleModel, error)

	// UserRoles
	CreateUserRole(userRole *model.UserRoleModel) error

	// General
	GetGeneralTiers() ([]*dto.TierModel, error)
	GetGeneralGenres(req dto.GeneralGenreRequest) ([]dto.GeneralGenreResponse, error)
	GetGeneralAuthors(req dto.GeneralAuthorRequest) ([]dto.GeneralAuthorResponse, error)

	// Ads
	CreateAds(ads *model.AdModel) error
	GetAds(id int64) (*dto.AdsResponse, error)
	GetAdsList(req dto.AdsListRequest) ([]dto.AdsResponse, int64, error)
	UpdateAds(ads *model.AdModel) error
	DeleteAds(id int64) error

	// Chapters
	CreateChapter(chapter *model.ChapterModel) error
	GetChapter(id int64) (*dto.ChapterResponse, error)
	ListChapters(req dto.ChapterListRequest) ([]dto.ChapterResponse, int64, error)
	UpdateChapter(chapter *model.ChapterModel) error
	DeleteChapter(id int64) error
	ActiveChapter(id int64, adminId int64) error

	// ChapterItems
	CreateChapterItem(item *model.ChapterItemModel) error
	GetChapterItem(id int64) (*dto.ChapterItemResponse, error)
	ListChapterItems(req dto.ChapterItemListRequest) ([]dto.ChapterItemResponse, int64, error)
	UpdateChapterItem(item *model.ChapterItemModel) error
	DeleteChapterItem(id int64) error

	CreateAuthor(author *model.AuthorModel) error
	GetAuthorById(id int64) (*model.AuthorModel, error)
	GetAuthors(req dto.RequestQueryFilter) ([]*model.AuthorModel, int64, error)
	UpdateAuthor(author *model.AuthorModel) (*model.AuthorModel, error)
	DeleteAuthorById(id int64) (*model.AuthorModel, error)
	GetAuthorsOfAComic(comicId int64) ([]dto.AuthorResponse, error)

	CreateRecommend(r *model.RecommendManagerModel) error
	GetRecommendById(id int64) (*model.RecommendManagerModel, error)
	GetRecommends(req dto.RequestQueryFilter) ([]*model.RecommendManagerModel, int64, error)
	DeleteRecommendById(id int64) (*model.RecommendManagerModel, error)
	UpdateRecommend(r *model.RecommendManagerModel) error
	CreateRecommendComic(r *model.RecommendComicModel) error
	DeleteRecommendComicById(comicId, recommendId int64) (*model.RecommendComicModel, error)
}

var _ Querier = (*Queries)(nil)
