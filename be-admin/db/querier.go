package db

type Querier interface {
	CreateUser(user *User) error
	GetUser(username string) (*User, error)
	DeleteUser(username string) error
	UpdateUser(user *User) error
}

var _ Querier = (*Queries)(nil)
