package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

type Claims struct {
	jwt.RegisteredClaims
	UserID int64 `json:"user_id"`
}

func NewJWTMaker(secretKey string) (Maker, error) {

	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d character", minSecretKeySize)
	}

	return &JWTMaker{secretKey}, nil
}
func (j *JWTMaker) CreateToken(userID int64, username string, role string, duration time.Duration) (string, *Payload, error) {

	payload, err := NewPayload(userID, username, role, duration)

	if err != nil {
		return "", payload, err
	}

	/*
		claims := &jwt.RegisteredClaims{
			ID:        payload.ID.String(),
			Issuer:    payload.Username,
			Subject:   payload.Role,
			ExpiresAt: jwt.NewNumericDate(payload.ExpiredAt),
			IssuedAt:  jwt.NewNumericDate(payload.IssuedAt),
		}
	*/
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        payload.ID.String(),
			Issuer:    payload.Username,
			Subject:   payload.Role,
			ExpiresAt: jwt.NewNumericDate(payload.ExpiredAt),
			IssuedAt:  jwt.NewNumericDate(payload.IssuedAt),
		},
		UserID: payload.UserID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString([]byte(j.secretKey))

	if err != nil {
		return "", payload, err
	}
	return jwtToken, payload, err
}

func (j *JWTMaker) VerifyToken(tokenID string) (*Payload, error) {

	// For checking SigningMethodHMAC is same or not base on previous usage in CreateToken func
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {

			return nil, ErrInvalidToken
		}
		return []byte(j.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(tokenID, &Claims{}, keyFunc, jwt.WithLeeway(5*time.Second))

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, jwt.ErrTokenExpired
		}

		return nil, ErrInvalidToken
	}

	if claim, ok := jwtToken.Claims.(*Claims); !ok {
		return nil, jwt.ErrTokenInvalidClaims
	} else {
		ID, err := uuid.Parse(claim.ID)

		if err != nil {
			return nil, err
		}

		payload := &Payload{
			ID:        ID,
			UserID:    claim.UserID,
			Username:  claim.Issuer,
			Role:      claim.Subject,
			IssuedAt:  claim.IssuedAt.Time,
			ExpiredAt: claim.ExpiresAt.Time,
		}

		return payload, nil
	}
}
