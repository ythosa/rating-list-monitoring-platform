package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/delivery/http/apierrors"
	"github.com/ythosa/rating-list-monitoring-platform-api/pkg/authorization"
)

const userCtx = "userID"

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
