package authorization

import (
	"errors"
	"fmt"
	"time"

	"github.com/ythosa/rating-list-monitoring-platform-api/pkg/config"

	"github.com/dgrijalva/jwt-go"
)

var ErrInvalidToken = errors.New("invalid token")

type JWTTokens struct {
	AccessToken  string
	RefreshToken string
}

type TokenClaims struct {
	jwt.StandardClaims
	UserID uint
}

func ParseToken(token string, tokenCfg config.JWTToken) (*TokenClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return tokenCfg.SigningKey, nil
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

func GenerateTokensFromPayload(userID uint, tokensConfig *config.AuthTokens) (*JWTTokens, error) {
	accessToken, err := GenerateTokenFromPayload(userID, tokensConfig.AccessToken)
	if err != nil {
		return nil, err
	}

	refreshToken, err := GenerateTokenFromPayload(userID, tokensConfig.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &JWTTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func GenerateTokenFromPayload(userID uint, tokenCfg config.JWTToken) (string, error) {
	tokenRaw := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenCfg.TTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: userID,
	})

	token, err := tokenRaw.SignedString(tokenCfg.SigningKey)
	if err != nil {
		return "", fmt.Errorf("error while signing token: %w", err)
	}

	return token, nil
}
