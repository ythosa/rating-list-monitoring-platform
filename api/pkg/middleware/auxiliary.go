package middleware

import (
	"errors"
	"strings"

	"github.com/ythosa/rating-list-monitoring-platform-api/pkg/apierrors"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	authorizationPrefix = "Bearer"
	refreshTokenHeader  = "RefreshToken"
)

const headerPartsCount = 2

func GetUserID(c *gin.Context) (uint, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("invalid user id")
	}

	userID, ok := id.(uint)
	if !ok {
		return 0, errors.New("invalid user id type")
	}

	return userID, nil
}

func GetAccessTokenFromRequest(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", apierrors.InvalidAuthorizationHeader
	}

	headersParts := strings.Split(header, " ")
	if len(headersParts) != headerPartsCount {
		return "", apierrors.InvalidAuthorizationHeader
	}

	if headersParts[0] != authorizationPrefix {
		return "", apierrors.InvalidAuthorizationHeader
	}

	accessToken := headersParts[1]
	if accessToken == "" {
		return "", apierrors.InvalidAuthorizationHeader
	}

	return accessToken, nil
}

func GetRefreshTokenFromRequest(c *gin.Context) (string, error) {
	refreshToken := c.GetHeader(refreshTokenHeader)
	if refreshToken == "" {
		return "", apierrors.InvalidRefreshToken
	}

	return refreshToken, nil
}
