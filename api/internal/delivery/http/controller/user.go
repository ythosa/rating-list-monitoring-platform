package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/delivery/http/apierrors"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/delivery/http/middleware"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/logging"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/service"
)

type UserImpl struct {
	validate    *validator.Validate
	userService service.User
	logger      *logging.Logger
}

func NewUserImpl(validate *validator.Validate, userService service.User) *UserImpl {
	return &UserImpl{
		validate:    validate,
		userService: userService,
		logger:      logging.NewLogger("user controller"),
	}
}

// GetUsername
// @tags user
// @summary returns user username
// @description returns user username by passing auth access token
// @accept json
// @produce json
// @security AccessTokenHeader
// @success 200 {object} dto.Username
// @failure 401 {object} apierrors.APIError
// @router /user/get_username [get].
func (u *UserImpl) GetUsername(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, apierrors.NewAPIError(err))

		return
	}

	username, err := u.userService.GetUsername(userID)
	if err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, apierrors.NewAPIError(err))

		return
	}

	c.JSON(http.StatusOK, username)
}

// GetProfile
// @tags user
// @summary returns user profile
// @description returns user username, firstname, lastname, middlename and snils
// @accept json
// @produce json
// @security AccessTokenHeader
// @success 200 {object} dto.UserProfile
// @failure 401 {object} apierrors.APIError
// @router /user/get_profile [get].
func (u *UserImpl) GetProfile(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, apierrors.NewAPIError(err))

		return
	}

	profile, err := u.userService.GetProfile(userID)
	if err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, apierrors.NewAPIError(err))

		return
	}

	c.JSON(http.StatusOK, profile)
}
