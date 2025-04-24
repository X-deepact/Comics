package middleware

import (
	"be-client/config"
	"be-client/util"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken     = errors.New("invalid token")
	ErrTokenExpired     = errors.New("token has expired")
	ErrMalformedToken   = errors.New("malformed token")
	ErrMissingToken     = errors.New("missing token")
	ErrInvalidSignature = errors.New("invalid token signature")
	Prefix              = "Bearer"
	RefreshTokenKeyName = "refreshToken"
	AccessTokenKeyName  = "accessToken"
	AuthorizationHeader = "Authorization"
	RefreshTokenHeader  = "RefreshToken"
)

type TokenPair struct {
	AccessToken  string    `json:"access_token,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	ExpiresAt    time.Time `json:"expires_at,omitempty"`
	TokenType    string    `json:"token_type,omitempty"`
}

type Claims struct {
	UserId    int64  `json:"user_id,omitempty"`
	UserName  string `json:"username,omitempty"`
	TokenType string `json:"token_type,omitempty"` // "access" or "refresh"
	jwt.RegisteredClaims
}

type AuthenHandler struct {
	config config.AuthenConfig
}

func NewAuthenHandler(config config.AuthenConfig) *AuthenHandler {
	return &AuthenHandler{
		config: config,
	}
}

func (a *AuthenHandler) GenerateAcessToken(userId int64, userName string) (*TokenPair, error) {
	accessToken, accessExp, err := a.generateToken(userId, userName, Prefix, a.config.AccessSecret, a.config.AccessTokenExpiration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}
	return &TokenPair{
		AccessToken: accessToken,
		ExpiresAt:   accessExp,
		TokenType:   Prefix,
	}, nil
}

func (a *AuthenHandler) GenerateTokenPair(userId int64, userName string) (*TokenPair, error) {
	accessToken, accessExp, err := a.generateToken(userId, userName, Prefix, a.config.AccessSecret, a.config.AccessTokenExpiration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, _, err := a.generateToken(userId, userName, Prefix, a.config.RefreshSecret, a.config.RefreshTokenExpiration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    accessExp,
		TokenType:    Prefix,
	}, nil
}

func (a *AuthenHandler) generateToken(
	userId int64,
	username string,
	tokenType string,
	secretKey string,
	expiration time.Duration,
) (string, time.Time, error) {
	now := time.Now()
	expiresAt := now.Add(expiration)

	claims := &Claims{
		UserId:    userId,
		UserName:  username,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Subject:   fmt.Sprintf("%d", userId),
			Issuer:    "client-side",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, expiresAt, nil
}

func (a *AuthenHandler) ValidateRefreshToken(tokenString string) (*Claims, error) {
	return a.ValidateToken(tokenString, a.config.RefreshSecret, Prefix)
}

func (a *AuthenHandler) ValidateToken(tokenString, secretToken, expectedType string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSignature
		}
		return []byte(secretToken), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, ErrMalformedToken
		}
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	if claims.TokenType != expectedType {
		return nil, fmt.Errorf("invalid token type: expected %s, got %s", expectedType, claims.TokenType)
	}

	return claims, nil
}

func (a *AuthenHandler) AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get(AuthorizationHeader)
		if auth == "" {
			return a.handleError(c, ErrMissingToken)
		}

		accessToken := strings.Split(auth, " ")
		if len(accessToken) != 2 || !strings.EqualFold(accessToken[0], Prefix) {
			return a.handleError(c, ErrMalformedToken)
		}

		claims, err := a.ValidateToken(accessToken[1], a.config.AccessSecret, Prefix)
		if err != nil {
			return a.handleError(c, err)
		}

		c.Locals("user", claims)
		c.Locals("user_id", claims.UserId)
		c.Locals("username", claims.UserName)

		return c.Next()
	}
}

func (a *AuthenHandler) handleError(c *fiber.Ctx, err error) error {
	status := fiber.StatusUnauthorized
	message := "Authentication failed"

	switch {
	case errors.Is(err, ErrMissingToken):
		status = fiber.StatusBadRequest
		message = "Missing authentication token"
	case errors.Is(err, ErrMalformedToken):
		status = fiber.StatusBadRequest
		message = "Malformed authentication token"
	case errors.Is(err, ErrTokenExpired):
		message = "Token has expired"
	case errors.Is(err, ErrInvalidSignature):
		message = "Invalid token signature"
	case errors.Is(err, ErrInvalidToken):
		message = "Invalid token"
	}
	slog.Error(fmt.Sprintf("Status error: %d, message: %s", status, message))
	return util.ResponseApi(c, nil, err)
}

func (a *AuthenHandler) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Get(AuthorizationHeader)
	if refreshToken == "" {
		return a.handleError(c, ErrMissingToken)
	}

	refreshToken = strings.TrimPrefix(refreshToken, Prefix)

	claims, err := a.ValidateToken(refreshToken, a.config.RefreshSecret, RefreshTokenKeyName)
	if err != nil {
		return a.handleError(c, err)
	}

	tokenPair, err := a.GenerateTokenPair(claims.UserId, claims.UserName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate new tokens",
		})
	}

	return c.JSON(tokenPair)
}

func (a *AuthenHandler) ExtractUserFromContext(c *fiber.Ctx) (*Claims, error) {
	user, ok := c.Locals("user").(*Claims)
	if !ok {
		return nil, errors.New("user not found in context")
	}
	return user, nil
}

func (r *AuthenHandler) GetUserNameFromContext(c *fiber.Ctx) (string, error) {
	userName := c.Locals("username")
	if userName == nil {
		return "", errors.New("user not found in context")
	}
	return userName.(string), nil
}

func (r *AuthenHandler) GetUserIdFromContext(c *fiber.Ctx) (int64, error) {
	userId := c.Locals("user_id")
	if userId == nil {
		return 0, errors.New("user not found in context")
	}
	return userId.(int64), nil
}

func (s *AuthenHandler) HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	h := hmac.New(sha256.New, []byte(s.config.PasswordSalt))

	h.Write([]byte(password))

	hash := h.Sum(nil)

	encodedHash := base64.StdEncoding.EncodeToString(hash)

	return encodedHash, nil
}
