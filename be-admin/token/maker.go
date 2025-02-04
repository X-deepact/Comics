package token

import "time"

type Maker interface {
	CreateToken(username string, role string, duration time.Duration) (string, *Payload, error)
	VerifyToken(tokenID string) (*Payload, error)
}
