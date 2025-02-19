package dto

type ListRequest struct {
	Page     int `form:"page" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=10,max=50"`
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}
type ListResponse struct {
	Pagination Pagination  `json:"pagination"`
	Data       interface{} `json:"data"`
}
