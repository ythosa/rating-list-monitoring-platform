package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/delivery/http/apierrors"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/delivery/http/middleware"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/logging"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/service"
	"net/http"
)

type UniversityImpl struct {
	validate          *validator.Validate
	universityService service.University
	logger            *logging.Logger
}

func NewUniversityImpl(validate *validator.Validate, universityService service.University) *UniversityImpl {
	return &UniversityImpl{
		validate:          validate,
		universityService: universityService,
		logger:            logging.NewLogger("university controller"),
	}
}

// GetAll
// @tags university
// @summary returns all universities
// @accept json
// @produce json
// @security AccessTokenHeader
// @success 200 {object} []rdto.University
// @failure 400 {object} apierrors.APIError
// @failure 401 {object} apierrors.APIError
// @router /university/get_all [get].
func (u *UniversityImpl) GetAll(c *gin.Context) {
	_, err := middleware.GetUserID(c)
	if err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, apierrors.NewAPIError(err))

		return
	}

	universities, err := u.universityService.GetAll()
	if err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.NewAPIError(err))

		return
	}

	c.JSON(http.StatusOK, universities)
}

// GetForUser
// @tags university
// @summary returns user universities
// @accept json
// @produce json
// @security AccessTokenHeader
// @success 200 {object} []rdto.University
// @failure 400 {object} apierrors.APIError
// @failure 401 {object} apierrors.APIError
// @router /university/get_for_user [get].
func (u *UniversityImpl) GetForUser(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, apierrors.NewAPIError(err))

		return
	}

	universities, err := u.universityService.GetForUser(userID)
	if err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.NewAPIError(err))

		return
	}

	c.JSON(http.StatusOK, universities)
}

// SetForUser
// @tags university
// @summary set universities to user
// @description receives university ids and sets it to user
// @accept json
// @produce json
// @security AccessTokenHeader
// @param payload body dto.IDs true "university ids"
// @success 200 "success"
// @failure 400 {object} apierrors.APIError
// @failure 401 {object} apierrors.APIError
// @router /university/set_for_user [post].
func (u *UniversityImpl) SetForUser(c *gin.Context) {
	var payload dto.IDs

	if err := c.BindJSON(&payload); err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.NewAPIError(err))

		return
	}

	userID, err := middleware.GetUserID(c)
	if err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, apierrors.NewAPIError(err))

		return
	}

	if err := u.universityService.SetForUser(userID, payload); err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.NewAPIError(err))

		return
	}

	c.Status(http.StatusOK)
}
