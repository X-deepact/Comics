package mapper

import (
	"be-client/config"
	"be-client/internal/dto/response"
	"be-client/util"
	"pkg-common/model"
)

type ChapterMapper interface {
	ToResponse(chapter *model.ChapterModel) response.ComicChapterResponse
	ToChapterDetailResponse(input *model.ChapterModel) response.ChapterDetailResponse
	ToChapterItemResponse(input *model.ChapterItemModel) response.ChapterItemResponse
	ToComicChapterTimeResponse(chapter map[string]interface{}) response.ComicChapterTimeResponse
	ToComicChapterResponse(chapter map[string]interface{}) response.ComicChapterResponse
}

type chapterMapper struct {
	config *config.Config
}

func NewChapterMapper(config *config.Config) ChapterMapper {
	return &chapterMapper{config}
}

func (m *chapterMapper) ToResponse(chapter *model.ChapterModel) response.ComicChapterResponse {
	if chapter == nil {
		return response.ComicChapterResponse{}
	}
	resp := response.ComicChapterResponse{
		Id:      chapter.Id,
		Name:    chapter.Name,
		IsCover: chapter.Cover,
		Number:  int64(chapter.Number),
	}
	if chapter.CreatedAt != nil {
		resp.CreatedAt = chapter.CreatedAt.UnixMilli()
	}
	return resp
}

func (m *chapterMapper) ToChapterDetailResponse(input *model.ChapterModel) response.ChapterDetailResponse {
	if input == nil {
		return response.ChapterDetailResponse{}
	}
	resp := response.ChapterDetailResponse{
		Id:      input.Id,
		Name:    input.Name,
		ComicId: input.ComicId,
		Cover:   input.Cover,
		Number:  input.Number,
		Active:  input.Active,
	}
	if input.CreatedAt != nil {
		resp.CreatedAt = input.CreatedAt.UnixMilli()
	}
	if input.ActiveFrom != nil {
		resp.ActiveFrom = input.ActiveFrom.UnixMilli()
	}
	return resp
}

func (m *chapterMapper) ToChapterItemResponse(input *model.ChapterItemModel) response.ChapterItemResponse {
	if input == nil {
		return response.ChapterItemResponse{}
	}
	resp := response.ChapterItemResponse{
		Page:   input.Page,
		Active: input.Active,
		Image:  util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.ChapterItemFilePath, input.Image),
	}
	if input.CreatedAt != nil {
		resp.CreatedAt = input.CreatedAt.UnixMilli()
	}
	if input.ActiveFrom != nil {
		resp.ActiveFrom = input.ActiveFrom.UnixMilli()
	}
	return resp
}

func (m *chapterMapper) ToComicChapterTimeResponse(chapter map[string]interface{}) response.ComicChapterTimeResponse {
	return response.ComicChapterTimeResponse{
		Id:        util.ToInt64(chapter["ch_id"]),
		Name:      util.GetStringOrEmpty(chapter["ch_name"]),
		IsCover:   util.ToInt64(chapter["ch_cover"]) == 1,
		Number:    util.ToInt64(chapter["ch_number"]),
		CreatedAt: util.ToTimePtr(chapter["ch_created_at"]),
	}
}

func (m *chapterMapper) ToComicChapterResponse(chapter map[string]interface{}) response.ComicChapterResponse {
	return response.ComicChapterResponse{
		Id:        util.ToInt64(chapter["ch_id"]),
		Name:      util.GetStringOrEmpty(chapter["ch_name"]),
		IsCover:   util.ToInt64(chapter["ch_cover"]) == 1,
		Number:    util.ToInt64(chapter["ch_number"]),
		CreatedAt: util.ToTimePtr(chapter["ch_created_at"]).UnixMilli(),
	}
}
