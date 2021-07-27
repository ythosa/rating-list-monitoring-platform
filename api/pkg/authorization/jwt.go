package authorization

import (
	"errors"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/config"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	AccessToken = iota
	RefreshToken
)

var (
	InvalidTokenError     = errors.New("invalid token")
	InvalidTokenTypeError = errors.New("invalid token type")
)

type TokenClaims struct {
	jwt.StandardClaims
	UserID uint8 `json:"user_id"`
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
		return nil, InvalidTokenTypeError
	}

	parsedToken, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return signingKey, nil
	})
	if err != nil {
		return nil, InvalidTokenError
	}

	claims, ok := parsedToken.Claims.(*TokenClaims)
	if !ok {
		return nil, InvalidTokenError
	}

	return claims, nil
}

func GenerateTokensFromPayload(userID uint8) (*dto.AuthorizationTokens, error) {
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

func GenerateAccessTokenFromPayload(userID uint8) (string, error) {
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
		return "", err
	}

	return token, nil
}

func GenerateRefreshTokenFromPayload(userID uint8) (string, error) {
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
		return "", err
	}

	return token, nil
}
