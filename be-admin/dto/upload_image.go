package dto

type UploadImageRequest struct {
	Image string `json:"image" binding:"required"`
}
