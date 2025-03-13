package dto

type ResponseMessage struct {
	Code       string      `json:"code,omitempty"`
	Msg        string      `json:"msg,omitempty"`
	Data       any         `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type SuccessResponse struct {
	Code string      `json:"code,omitempty" example:"SUCCESS"`
	Msg  string      `json:"msg,omitempty" example:"success"`
	Data interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code string `json:"code,omitempty" example:"ERROR"`
	Msg  string `json:"msg,omitempty"`
}

type ListSuccessResponse struct {
	SuccessResponse
	Pagination Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Page     int `json:"page" example:"1"`
	PageSize int `json:"page_size" example:"10"`
	Total    int `json:"total" example:"100"`
}
