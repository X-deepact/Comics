package config

const (
	ADMIN = "admin"
	USER  = "user"
)

func IsSupportedRoles(role string) bool {
	switch role {
	case ADMIN, USER:
		return true

	}
	return false
}
