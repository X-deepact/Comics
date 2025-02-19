package dto

type RecommendCreateRequest struct {
	Id         int64
	Title      string
	Cover      string
	Position   int
	ActiveFrom int64
	ActiveTo   int64
}

type RecommendResponse struct {
	Id         int64
	Title      string
	Cover      string
	Position   int
	ActiveFrom string
	ActiveTo   string
	CreatedAt  string
	CreatedBy  int64
	UpdatedAt  string
	UpdatedBy  int64
}

type RecommendUpdateRequest struct {
	Title      string
	Cover      string
	Position   int
	ActiveFrom string
	ActiveTo   string
}
