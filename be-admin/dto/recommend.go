package dto

type RecommendCreateRequest struct {
	Title      string `json:"title,omitempty"`
	Cover      string `json:"cover,omitempty"`
	Position   int    `json:"position,omitempty"`
	ActiveFrom int64  `json:"active_from,omitempty"`
	ActiveTo   int64  `json:"active_to,omitempty"`
}

type RecommendResponse struct {
	Id         int64  `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Cover      string `json:"cover,omitempty"`
	Position   int    `json:"position,omitempty"`
	ActiveFrom string `json:"active_from,omitempty"`
	ActiveTo   string `json:"active_to,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	CreatedBy  int64  `json:"created_by,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	UpdatedBy  int64  `json:"updated_by,omitempty"`
}

type RecommendUpdateRequest struct {
	Id         int64  `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Cover      string `json:"cover,omitempty"`
	Position   int    `json:"position,omitempty"`
	ActiveFrom string `json:"active_from,omitempty"`
	ActiveTo   string `json:"active_to,omitempty"`
}
