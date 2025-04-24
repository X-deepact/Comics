package response

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid request parameters"`
}

// Common error responses
var (
	ErrInvalidRequest = &ErrorResponse{Error: "Invalid request parameters"}
	ErrNotFound       = &ErrorResponse{Error: "Resource not found"}
	ErrUnauthorized   = &ErrorResponse{Error: "Unauthorized access"}
	ErrServerError    = &ErrorResponse{Error: "Internal server error"}
)

func ErrorBind(err error) *ErrorResponse {
	return &ErrorResponse{Error: err.Error()}
}
