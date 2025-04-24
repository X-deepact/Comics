package mapper

import (
	"be-client/config"
	"be-client/internal/dto/response"
	"be-client/util"
	"pkg-common/model"
)

type AuthorMapper interface {
	ToAuthorComicResponse(input []*model.AuthorModel) []response.ComicAuthorResponse
	ToComicAuthorTimeResponse(author map[string]interface{}) response.ComicAuthorTimeResponse
}

type authorMapper struct {
	config *config.Config
}

func NewAuthorMapper(config *config.Config) AuthorMapper {
	return &authorMapper{config}
}

func (m *authorMapper) ToAuthorComicResponse(input []*model.AuthorModel) []response.ComicAuthorResponse {
	if len(input) == 0 {
		return nil
	}
	resp := make([]response.ComicAuthorResponse, len(input))
	for i, author := range input {
		resp[i] = response.ComicAuthorResponse{
			Id:        author.Id,
			Name:      author.Name,
			Biography: author.Biography,
		}
		if author.CreatedAt != nil {
			resp[i].CreatedAt = author.CreatedAt.UnixMilli()
		}
		if author.UpdatedAt != nil {
			resp[i].UpdatedAt = author.UpdatedAt.UnixMilli()
		}
		if author.BirthDay != nil {
			resp[i].BirthDay = author.BirthDay.UnixMilli()
		}
	}
	return resp
}

func (m *authorMapper) ToComicAuthorTimeResponse(author map[string]interface{}) response.ComicAuthorTimeResponse {
	return response.ComicAuthorTimeResponse{
		Id:        util.ToInt64(author["aut_id"]),
		Name:      util.GetStringOrEmpty(author["aut_name"]),
		Biography: util.GetStringOrEmpty(author["aut_biography"]),
		BirthDay:  util.ToTimePtr(author["aut_birthday"]),
		CreatedAt: util.ToTimePtr(author["aut_created_at"]),
		UpdatedAt: util.ToTimePtr(author["aut_updateted_at"]),
	}
}
