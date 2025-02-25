package token

import "time"

type Maker interface {
	CreateToken(userID int64, username string, role string, duration time.Duration) (string, *Payload, error)
	VerifyToken(tokenID string) (*Payload, error)
}
