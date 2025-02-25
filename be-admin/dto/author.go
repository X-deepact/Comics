package dto

type AuthorRequest struct {
	Name      string `json:"name" binding:"required"`
	BirthDay  string `json:"birth_day"`
	Biography string `json:"biography"`
}

type AuthorResponse struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Biography     string `json:"biography,omitempty"`
	BirthDay      string `json:"birth_day,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	CreatedByName string `json:"created_by_name,omitempty"`
	UpdatedByName string `json:"updated_by_name,omitempty"`
}

type AuthorUpdateRequest struct {
	Id        int64  `json:"id" binding:"required"`
	Name      string `json:"name"`
	BirthDay  string `json:"birth_day"`
	Biography string `json:"biography"`
}
