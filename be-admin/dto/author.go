package dto

type AuthorRequest struct {
	Name      string `json:"name" binding:"required"`
	BirthDay  string `json:"birth_day"`
	Biography string `json:"biography"`
}

type AuthorResponse struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
	BirthDay  string `json:"birth_day"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	CreatedBy int64  `json:"created_by"`
	UpdatedBy int64  `json:"updated_by"`
}

type AuthorUpdateRequest struct {
	Name      string `json:"name"`
	BirthDay  string `json:"birth_day"`
	Biography string `json:"biography"`
}
