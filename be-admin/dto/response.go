package dto

type ResponseMessage struct {
	Code       string      `json:"code,omitempty"`
	Msg        string      `json:"msg,omitempty"`
	Data       any         `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}
