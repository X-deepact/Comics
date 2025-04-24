package mapper

import (
	"be-client/config"
	"be-client/internal/dto/response"
	"be-client/util"
	"pkg-common/model"
)

type ProfileMapper interface {
	ToResponse(req *model.ProfileModel) *response.ProfileResponse
	ToPrfileTierResponse(req *model.TierModel) *response.ProfileTierResponse
}

type profileMapper struct {
	config *config.Config
}

func NewProfileMapper(config *config.Config) ProfileMapper {
	return &profileMapper{config: config}
}

func (m *profileMapper) ToResponse(req *model.ProfileModel) *response.ProfileResponse {
	resp := &response.ProfileResponse{
		Id:          req.Id,
		Active:      req.Active,
		DisplayName: req.DisplayName,
		Description: req.Description,
		Avatar:      util.GetFileUrl(m.config.FileStorage.EndPoint, m.config.FileStorage.BucketName, m.config.FileStorage.AvatarFilePath, req.Avatar),
		UserId:      req.UserId,
	}
	if req.CreatedAt != nil {
		resp.CreatedAt = req.CreatedAt.UnixMilli()
	}
	if req.UpdatedAt != nil {
		resp.UpdatedAt = req.UpdatedAt.UnixMilli()
	}
	return resp
}

func (m *profileMapper) ToPrfileTierResponse(req *model.TierModel) *response.ProfileTierResponse {
	return &response.ProfileTierResponse{
		Id:        req.Id,
		Code:      req.Code,
		CreatedAt: req.CreatedAt.UnixMilli(),
	}
}
