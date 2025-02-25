package config

const (
	ADMIN = "Admin"
	USER  = "User"
)

func IsSupportedRoles(role string) bool {
	switch role {
	case ADMIN, USER:
		return true

	}
	return false
}
