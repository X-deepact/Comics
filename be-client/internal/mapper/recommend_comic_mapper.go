package mapper

import (
	"be-client/config"
	"be-client/internal/dto/response"
	"be-client/util"
	"pkg-common/model"
	"sort"
)

type RecommendComicMapper interface {
	ToResponseList(recommends *[]model.RecommendManagerModel, recommendComicList *[]map[string]interface{}, comicChapterList *[]map[string]interface{}, chapterMapper ChapterMapper) []response.RecommendComicResponse
	ToResponseList1(recommends *[]model.RecommendManagerModel, recommendComicList *[]map[string]interface{}, comicChapterList *[]map[string]interface{}, chapterMapper ChapterMapper) []response.RecommendComicResponse
	MapPositionToType(position int) string
}

type recommendComicMapper struct {
	config *config.Config
}

func NewRecommendComicMapper(config *config.Config) RecommendComicMapper {
	return &recommendComicMapper{config}
}

func (m *recommendComicMapper) ToResponseList(
	recommends *[]model.RecommendManagerModel,
	recommendComicList *[]map[string]interface{},
	comicChapterList *[]map[string]interface{},
	chapterMapper ChapterMapper,
) []response.RecommendComicResponse {
	var responses []response.RecommendComicResponse

	recommendManagerMap := make(map[int64]*response.RecommendComicResponse)
	positionMap := make(map[int64]int64)
	comicMap := make(map[int64]*response.ComicWithChapterTimeResponse)
	type tempResponse struct {
		Position int64
		Response response.RecommendComicResponse
	}

	for _, recommend := range *recommends {
		rmId := recommend.Id
		rmPosition := recommend.Position
		rmPositionString := m.MapPositionToType(int(rmPosition))
		if _, exists := recommendManagerMap[rmId]; !exists {
			recommendManagerMap[rmId] = &response.RecommendComicResponse{
				Id:         rmId,
				Title:      recommend.Title,
				Cover:      util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.RecommendFilePath, recommend.Cover),
				Position:   rmPositionString,
				ActiveFrom: recommend.ActiveFrom,
				ActiveTo:   recommend.ActiveTo,
				CreatedAt:  recommend.CreatedAt,
				UpdatedAt:  recommend.UpdatedAt,
				Comics:     []response.ComicWithChapterTimeResponse{},
			}
			positionMap[rmId] = util.ToInt64(rmPosition)
		}
	}

	for _, comic := range *recommendComicList {
		rmId := util.ToInt64(comic["rm_id"])
		cmId := util.ToInt64(comic["cm_id"])

		if recommend, exists := recommendManagerMap[rmId]; exists {
			comicResponse := response.ComicWithChapterTimeResponse{
				Id:          cmId,
				Name:        util.GetStringOrEmpty(comic["cm_name"]),
				Code:        util.GetStringOrEmpty(comic["cm_code"]),
				Language:    util.GetStringOrEmpty(comic["cm_language"]),
				Audience:    util.GetStringOrEmpty(comic["cm_audience"]),
				Description: util.GetStringOrEmpty(comic["cm_description"]),
				Cover:       util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.ComicCoverFilePath, util.GetStringOrEmpty(comic["cm_cover"])),
				Active:      util.ToInt64(comic["cm_active"]) == 1,
				CreatedAt:   util.ToTimePtr(comic["cm_created_at"]),
				UpdatedAt:   util.ToTimePtr(comic["cm_updated_at"]),
				Chapters:    []response.ComicChapterTimeResponse{},
			}

			comicMap[cmId] = &comicResponse
			recommend.Comics = append(recommend.Comics, comicResponse)
		}
	}

	for _, chapter := range *comicChapterList {
		cmId := util.ToInt64(chapter["cm_id"])

		if comic, exists := comicMap[cmId]; exists {
			comic.Chapters = append(comic.Chapters, chapterMapper.ToComicChapterTimeResponse(chapter))
		}
	}

	for _, recommend := range recommendManagerMap {
		for i, comic := range recommend.Comics {
			if updatedComic, exists := comicMap[comic.Id]; exists {
				recommend.Comics[i] = *updatedComic
			}
		}
	}

	var tempResponses []tempResponse
	for rmId, recommend := range recommendManagerMap {
		tempResponses = append(tempResponses, tempResponse{
			Position: positionMap[rmId],
			Response: *recommend,
		})
	}

	sort.Slice(tempResponses, func(i, j int) bool {
		return tempResponses[i].Position < tempResponses[j].Position
	})

	for _, temp := range tempResponses {
		responses = append(responses, temp.Response)
	}

	return responses
}

func (r *recommendComicMapper) MapPositionToType(position int) string {
	switch position {
	case 1:
		return "TOPING"
	case 2:
		return "RECOMMEND_PRODUCTS"
	case 3:
		return "RECOMMEND_MASTERPIECES"
	case 4:
		return "FASTEST_GROWING"
	case 5:
		return "TESTING_NEW_PRODUCTS"
	default:
		return ""
	}
}

func (m *recommendComicMapper) ToResponseList1(
	recommends *[]model.RecommendManagerModel,
	recommendComicList *[]map[string]interface{},
	comicChapterList *[]map[string]interface{},
	chapterMapper ChapterMapper,
) []response.RecommendComicResponse {
	var responses []response.RecommendComicResponse

	positionComicsMap := make(map[int64][]response.ComicWithChapterTimeResponse)
	chapterMap := make(map[int64][]response.ComicChapterTimeResponse)
	recommendManagerMap := make(map[int64]*response.RecommendComicResponse)
	for _, recommend := range *recommends {
		rmId := recommend.Id
		rmPosition := recommend.Position
		rmPositionString := m.MapPositionToType(int(rmPosition))
		if _, exists := recommendManagerMap[rmId]; !exists {
			recommendManagerMap[rmId] = &response.RecommendComicResponse{
				Id:         rmId,
				Title:      recommend.Title,
				Cover:      util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.RecommendFilePath, recommend.Cover),
				Position:   rmPositionString,
				ActiveFrom: recommend.ActiveFrom,
				ActiveTo:   recommend.ActiveTo,
				CreatedAt:  recommend.CreatedAt,
				UpdatedAt:  recommend.UpdatedAt,
				Comics:     []response.ComicWithChapterTimeResponse{},
			}

		}
	}

	for _, chapter := range *comicChapterList {
		cmId := util.ToInt64(chapter["cm_id"])
		cha := chapterMapper.ToComicChapterTimeResponse(chapter)
		if v, ok := chapterMap[cmId]; ok {
			v = append(v, cha)
		} else {
			chapterMap[cmId] = []response.ComicChapterTimeResponse{cha}
		}
	}

	for _, rc := range *recommendComicList {
		cmId := util.ToInt64(rc["cm_id"])
		position := util.ToInt64(rc["rm_position"])
		comic := response.ComicWithChapterTimeResponse{
			Id:          cmId,
			Name:        util.GetStringOrEmpty(rc["cm_name"]),
			Code:        util.GetStringOrEmpty(rc["cm_code"]),
			Language:    util.GetStringOrEmpty(rc["cm_language"]),
			Audience:    util.GetStringOrEmpty(rc["cm_audience"]),
			Description: util.GetStringOrEmpty(rc["cm_description"]),
			Cover:       util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.ComicCoverFilePath, util.GetStringOrEmpty(rc["cm_cover"])),
			Active:      util.ToInt64(rc["cm_active"]) == 1,
			CreatedAt:   util.ToTimePtr(rc["cm_created_at"]),
			UpdatedAt:   util.ToTimePtr(rc["cm_updated_at"]),
			Chapters:    []response.ComicChapterTimeResponse{},
		}
		if chapters, ok := chapterMap[cmId]; ok {
			comic.Chapters = chapters
		}
		if v, ok := positionComicsMap[position]; ok {
			// each type recommendation limit to 10
			if len(v) >= 10 {
				continue
			}
			positionComicsMap[position] = append(v, comic)
		} else {
			positionComicsMap[position] = []response.ComicWithChapterTimeResponse{comic}
		}
	}

	type tempResponse struct {
		Position int64
		Response response.RecommendComicResponse
	}
	var tempResponses []tempResponse

	for p, v := range positionComicsMap {
		var tm response.RecommendComicResponse
		positionStr := ""

		switch p {
		case 1:
			positionStr = "TOPING"
		case 2:
			positionStr = "RECOMMEND_PRODUCTS"
		case 3:
			positionStr = "COMPLETED_MASTERPIECES"
		case 4:
			positionStr = "FASTEST_GROWING"
		case 5:
			positionStr = "TESTING_NEW_PRODUCTS"
		}
		tm.Position = positionStr
		tm.Comics = v
		tempResponses = append(tempResponses, tempResponse{p, tm})
	}

	sort.Slice(tempResponses, func(i, j int) bool {
		return tempResponses[i].Position < tempResponses[j].Position
	})

	for _, temp := range tempResponses {
		responses = append(responses, temp.Response)
	}

	return responses
}
