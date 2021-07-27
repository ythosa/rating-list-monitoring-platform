package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/delivery/http/apierrors"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/pkg/authorization"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	refreshTokenHeader  = "RefreshToken"
)

const userCtx = "userID"

const headerPartsCount = 2

func UserIdentity(c *gin.Context) {
	accessToken, err := GetAccessTokenFromRequest(c)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, apierrors.NewAPIError(err))

		return
	}

	tokenClaims, err := authorization.ParseToken(accessToken, authorization.AccessToken)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, apierrors.NewAPIError(apierrors.InvalidAuthorizationHeader))

		return
	}

	c.Set(userCtx, tokenClaims.UserID)
}

func GetUserID(c *gin.Context) (uint8, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("invalid user id")
	}

	userID, ok := id.(uint8)
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
