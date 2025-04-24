package mapper

import (
	"be-client/config"
	"be-client/internal/dto/response"
	"be-client/util"
	"pkg-common/model"
	"sort"
)

type ComicMapper interface {
	ToResponse(chapterMapper ChapterMapper, authorMapper AuthorMapper, data *[]model.ComicModel, comicChapters *[]map[string]interface{}) []response.ComicWithChapterResponse
	FromModelToComicResponse(data *model.ComicModel) response.ComicResponse
	ToComicDetailResponse(data *model.ComicModel) response.ComicDetailResponse
	ToComicSearchResponse(chapterMapper ChapterMapper, authorMapper AuthorMapper, data *[]model.ComicModel, comicAuthors *[]map[string]interface{}, comicChapters *[]map[string]interface{}) []response.ComicWithChapterAndAuthorResponse
	ToComicRecentResponse(chapterMapper ChapterMapper, data *[]model.ComicModel, lastChapterInfo *[]map[string]interface{}, comicChapters *[]map[string]interface{}) []response.ComicRecent
	ToComicAndDramaSearchResponse(chapterMapper ChapterMapper, authorMapper AuthorMapper, data *[]map[string]interface{}, comicAuthors *[]map[string]interface{}, comicChapters *[]map[string]interface{}, lastEpNum *[]map[string]interface{}) []response.ComicAndDramaSearchResponse
}
type comicMapper struct {
	config *config.Config
}

func NewComicMapper(config *config.Config) ComicMapper {
	return &comicMapper{config}
}

func (m *comicMapper) ToResponse(
	chapterMapper ChapterMapper,
	authorMapper AuthorMapper,
	data *[]model.ComicModel,
	comicChapters *[]map[string]interface{},
) []response.ComicWithChapterResponse {
	if len(*data) == 0 {
		return []response.ComicWithChapterResponse{}
	}

	comicResponsesMap := make(map[int64]*response.ComicWithChapterResponse)
	comicResponses := make([]*response.ComicWithChapterResponse, 0, len(*data))

	for _, comic := range *data {
		cmId := comic.Id

		if _, exists := comicResponsesMap[cmId]; !exists {
			comicResponse := &response.ComicWithChapterResponse{
				Id:          cmId,
				Name:        comic.Name,
				Code:        comic.Code,
				Language:    comic.Lang,
				Audience:    comic.Audience,
				Description: comic.Description,
				Cover:       util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.ComicCoverFilePath, comic.Cover),
				Active:      comic.Active,
				CreatedAt:   comic.CreatedAt.UnixMilli(),
				UpdatedAt:   comic.UpdatedAt.UnixMilli(),
				Chapters:    []response.ComicChapterResponse{},
			}

			comicResponsesMap[cmId] = comicResponse
			comicResponses = append(comicResponses, comicResponse)
		}
	}

	for _, chapter := range *comicChapters {
		cmId := util.ToInt64(chapter["cm_id"])
		if comic, exists := comicResponsesMap[cmId]; exists {
			comic.Chapters = append(comic.Chapters, chapterMapper.ToComicChapterResponse(chapter))
		}
	}

	finalComicResponses := make([]response.ComicWithChapterResponse, len(comicResponses))
	for i, comic := range comicResponses {
		finalComicResponses[i] = *comic
	}

	return finalComicResponses
}

func (m *comicMapper) FromModelToComicResponse(data *model.ComicModel) response.ComicResponse {
	comicResponse := response.ComicResponse{
		Id:          data.Id,
		Name:        data.Name,
		Code:        data.Code,
		Language:    data.Lang,
		Audience:    data.Audience,
		Description: data.Description,
		Cover:       util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.ComicCoverFilePath, data.Cover),
		Active:      data.Active,
	}
	if data.CreatedAt != nil {
		comicResponse.CreatedAt = data.CreatedAt.UnixMilli()
	}
	if data.UpdatedAt != nil {
		comicResponse.UpdatedAt = data.UpdatedAt.UnixMilli()
	}
	return comicResponse
}

func (m *comicMapper) ToComicDetailResponse(data *model.ComicModel) response.ComicDetailResponse {
	if data == nil {
		return response.ComicDetailResponse{}
	}
	comicResponse := response.ComicDetailResponse{
		Id:          data.Id,
		Name:        data.Name,
		Code:        data.Code,
		Language:    data.Lang,
		Audience:    data.Audience,
		Description: data.Description,
		Cover:       util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.ComicCoverFilePath, data.Cover),
		Active:      data.Active,
		Status:      data.Status,
	}
	if data.CreatedAt != nil {
		comicResponse.CreatedAt = data.CreatedAt.UnixMilli()
	}
	if data.UpdatedAt != nil {
		comicResponse.UpdatedAt = data.UpdatedAt.UnixMilli()
	}
	return comicResponse
}

func (m *comicMapper) ToComicSearchResponse(
	chapterMapper ChapterMapper,
	authorMapper AuthorMapper,
	data *[]model.ComicModel,
	comicAuthors *[]map[string]interface{},
	comicChapters *[]map[string]interface{},
) []response.ComicWithChapterAndAuthorResponse {
	if len(*data) == 0 {
		return []response.ComicWithChapterAndAuthorResponse{}
	}

	comicResponsesMap := make(map[int64]*response.ComicWithChapterAndAuthorResponse)
	comicResponses := make([]*response.ComicWithChapterAndAuthorResponse, 0, len(*data))

	for _, comic := range *data {
		cmId := comic.Id

		if _, exists := comicResponsesMap[cmId]; !exists {
			comicResponse := &response.ComicWithChapterAndAuthorResponse{
				Id:          cmId,
				Name:        comic.Name,
				Code:        comic.Code,
				Language:    comic.Lang,
				Audience:    comic.Audience,
				Description: comic.Description,
				Cover:       util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.ComicCoverFilePath, comic.Cover),
				Active:      comic.Active,
				CreatedAt:   comic.CreatedAt,
				UpdatedAt:   comic.UpdatedAt,
				Chapters:    []response.ComicChapterTimeResponse{},
				Authors:     []response.ComicAuthorTimeResponse{},
			}

			comicResponsesMap[cmId] = comicResponse
			comicResponses = append(comicResponses, comicResponse)
		}
	}

	for _, chapter := range *comicChapters {
		cmId := util.ToInt64(chapter["cm_id"])
		if comic, exists := comicResponsesMap[cmId]; exists {
			comic.Chapters = append(comic.Chapters, chapterMapper.ToComicChapterTimeResponse(chapter))
		}
	}

	for _, author := range *comicAuthors {
		cmId := util.ToInt64(author["cm_id"])
		if comic, exists := comicResponsesMap[cmId]; exists {
			comic.Authors = append(comic.Authors, authorMapper.ToComicAuthorTimeResponse(author))
		}
	}

	finalComicResponses := make([]response.ComicWithChapterAndAuthorResponse, len(comicResponses))
	for i, comic := range comicResponses {
		finalComicResponses[i] = *comic
	}

	return finalComicResponses
}

func (m *comicMapper) ToComicRecentResponse(chapterMapper ChapterMapper, data *[]model.ComicModel, lastChapterInfo *[]map[string]interface{}, comicChapters *[]map[string]interface{}) []response.ComicRecent {
	if len(*data) == 0 {
		return []response.ComicRecent{}
	}

	comicResponsesMap := make(map[int64]*response.ComicRecent)
	comicResponses := []response.ComicRecent{}

	for _, comic := range *data {
		cmId := comic.Id

		if _, exists := comicResponsesMap[cmId]; !exists {
			comicResponsesMap[cmId] = &response.ComicRecent{
				Id:          cmId,
				Name:        comic.Name,
				Code:        comic.Code,
				Language:    comic.Lang,
				Audience:    comic.Audience,
				Description: comic.Description,
				Cover:       util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.ComicCoverFilePath, comic.Cover),
				Active:      comic.Active,
				Chapters:    []response.ComicChapterTimeResponse{},
			}
		}
	}

	for _, chapter := range *comicChapters {
		cmId := util.ToInt64(chapter["cm_id"])

		if comic, exists := comicResponsesMap[cmId]; exists {
			comic.Chapters = append(comic.Chapters, chapterMapper.ToComicChapterTimeResponse(chapter))
		}
	}

	for _, chapter := range *lastChapterInfo {
		cmId := util.ToInt64(chapter["cm_id"])
		last_date := util.ToTimePtr(chapter["last_created_date"])

		if chapter, exists := comicResponsesMap[cmId]; exists {
			chapter.LastChapterCreDate = last_date
		}
	}

	for _, comic := range comicResponsesMap {
		comicResponses = append(comicResponses, *comic)
	}

	sort.Slice(comicResponses, func(i, j int) bool {
		if comicResponses[i].LastChapterCreDate != nil && comicResponses[j].LastChapterCreDate != nil {
			if (*comicResponses[i].LastChapterCreDate).Equal(*comicResponses[j].LastChapterCreDate) {
				return comicResponses[i].Id < comicResponses[j].Id
			}
			return (*comicResponses[i].LastChapterCreDate).After(*comicResponses[j].LastChapterCreDate)
		}
		return comicResponses[i].LastChapterCreDate != nil
	})

	return comicResponses
}

func (m *comicMapper) ToComicAndDramaSearchResponse(
	chapterMapper ChapterMapper,
	authorMapper AuthorMapper,
	data *[]map[string]interface{},
	comicAuthors *[]map[string]interface{},
	comicChapters *[]map[string]interface{},
	lastEpNum *[]map[string]interface{},
) []response.ComicAndDramaSearchResponse {
	if len(*data) == 0 {
		return []response.ComicAndDramaSearchResponse{}
	}

	lastNumMap := make(map[int64]int64)
	for _, ep := range *lastEpNum {
		dramaID := util.ToInt64(ep["drama_id"])
		lastNum := util.ToInt64(ep["number"])
		lastNumMap[dramaID] = lastNum
	}

	comicResponsesMap := make(map[int64]*response.ComicAndDramaSearchResponse)
	comicResponses := make([]*response.ComicAndDramaSearchResponse, 0, len(*data))

	for _, comic := range *data {
		cmId := util.ToInt64(comic["id"])
		comicType := util.GetStringOrEmpty(comic["type"])

		if _, exists := comicResponsesMap[cmId]; !exists {
			comicResponse := &response.ComicAndDramaSearchResponse{
				Id:          cmId,
				Type:        comicType,
				Name:        util.GetStringOrEmpty(comic["name"]),
				Code:        util.GetStringOrEmpty(comic["code"]),
				Language:    util.GetStringOrEmpty(comic["lang"]),
				Audience:    util.GetStringOrEmpty(comic["audience"]),
				Description: util.GetStringOrEmpty(comic["description"]),
				Cover:       util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.ComicCoverFilePath, util.GetStringOrEmpty(comic["cover"])),
				Active:      util.ToInt64(comic["active"]) == 1,
				CreatedAt:   util.ToTimePtr(comic["created_at"]),
				UpdatedAt:   util.ToTimePtr(comic["updated_at"]),
				Chapters:    []response.ComicChapterTimeResponse{},
				Authors:     []response.ComicAuthorTimeResponse{},
				ReleaseDate: util.ToTimePtr(comic["release_date"]),
				LatestEpisodeNumber: func() int64 {
					if comicType == "drama" {
						return lastNumMap[cmId]
					}
					return 0
				}(),
				Thumb: util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.DramaThumbFilePath, util.GetStringOrEmpty(comic["thumb"])),
			}

			comicResponsesMap[cmId] = comicResponse
			comicResponses = append(comicResponses, comicResponse)
		}
	}

	for _, chapter := range *comicChapters {
		cmId := util.ToInt64(chapter["cm_id"])
		if comic, exists := comicResponsesMap[cmId]; exists {
			if comic.Type == "comic" {
				comic.Chapters = append(comic.Chapters, chapterMapper.ToComicChapterTimeResponse(chapter))
			}

		}
	}

	for _, author := range *comicAuthors {
		cmId := util.ToInt64(author["cm_id"])
		if comic, exists := comicResponsesMap[cmId]; exists {
			if comic.Type == "comic" {
				comic.Authors = append(comic.Authors, authorMapper.ToComicAuthorTimeResponse(author))
			}

		}
	}

	finalComicResponses := make([]response.ComicAndDramaSearchResponse, len(comicResponses))
	for i, comic := range comicResponses {
		finalComicResponses[i] = *comic
	}

	return finalComicResponses
}
