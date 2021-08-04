package authorization

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/config"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/dto"
)

const (
	AccessToken = iota
	RefreshToken
)

var (
	ErrInvalidToken     = errors.New("invalid token")
	ErrInvalidTokenType = errors.New("invalid token type")
)

type TokenClaims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
}

func ParseToken(token string, tokenType int) (*TokenClaims, error) {
	tokensCfg := config.Get().Authorization

	var signingKey []byte

	switch tokenType {
	case AccessToken:
		signingKey = tokensCfg.AccessToken.SigningKey
	case RefreshToken:
		signingKey = tokensCfg.RefreshToken.SigningKey
	default:
		return nil, ErrInvalidTokenType
	}

	parsedToken, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return signingKey, nil
	})
	if err != nil {
		return nil, ErrInvalidToken
	}

	claims, ok := parsedToken.Claims.(*TokenClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func GenerateTokensFromPayload(userID uint) (*dto.AuthorizationTokens, error) {
	accessToken, err := GenerateAccessTokenFromPayload(userID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := GenerateRefreshTokenFromPayload(userID)
	if err != nil {
		return nil, err
	}

	return &dto.AuthorizationTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func GenerateAccessTokenFromPayload(userID uint) (string, error) {
	cfg := config.Get().Authorization.AccessToken

	tokenRaw := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(cfg.TTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: userID,
	})

	token, err := tokenRaw.SignedString(cfg.SigningKey)
	if err != nil {
		return "", fmt.Errorf("error while signing token: %w", err)
	}

	return token, nil
}

func GenerateRefreshTokenFromPayload(userID uint) (string, error) {
	cfg := config.Get().Authorization.RefreshToken

	tokenRaw := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(cfg.TTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: userID,
	})

	token, err := tokenRaw.SignedString(cfg.SigningKey)
	if err != nil {
		return "", fmt.Errorf("error while signing token: %w", err)
	}

	return token, nil
}
