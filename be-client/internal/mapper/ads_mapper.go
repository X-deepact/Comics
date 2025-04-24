package mapper

import (
	"be-client/config"
	"be-client/internal/dto/response"
	"be-client/util"
)

type AdsMapper interface {
	ToAdsComicResponse(data *[]map[string]interface{}) []response.AdsComicResponse
}
type adsMapper struct {
	config *config.Config
}

func NewAdsMapper(config *config.Config) AdsMapper {
	return &adsMapper{config}
}

func (m *adsMapper) ToAdsComicResponse(data *[]map[string]interface{}) []response.AdsComicResponse {
	if len(*data) == 0 {
		return nil
	}

	adsResponses := make([]response.AdsComicResponse, len(*data))

	for i, ads := range *data {
		adsResponses[i] = response.AdsComicResponse{
			A_Id:          util.ToInt64(ads["ads_id"]),
			A_Title:       util.GetStringOrEmpty(ads["ads_title"]),
			A_Image:       util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.AdsFilePath, util.GetStringOrEmpty(ads["ads_image"])),
			A_Type:        util.GetStringOrEmpty(ads["ads_type"]),
			A_DirectURL:   util.GetStringOrEmpty(ads["ads_direct_url"]),
			A_ActiveFrom:  util.ToTimePtr(ads["ads_active_from"]),
			A_ActiveTo:    util.ToTimePtr(ads["ads_active_to"]),
			C_Id:          util.ToInt64(ads["comic_id"]),
			C_Name:        util.GetStringOrEmpty(ads["comic_name"]),
			C_Code:        util.GetStringOrEmpty(ads["comic_code"]),
			C_Language:    util.GetStringOrEmpty(ads["comic_language"]),
			C_Audience:    util.GetStringOrEmpty(ads["comic_audience"]),
			C_Description: util.GetStringOrEmpty(ads["comic_description"]),
			C_Cover:       util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.ComicCoverFilePath, util.GetStringOrEmpty(ads["comic_cover"])),
			C_Active:      util.ToInt64(ads["comic_active"]) == 1,
		}
	}

	return adsResponses
}
