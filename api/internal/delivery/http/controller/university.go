package controller

import (
	"net/http"
	"strconv"

	"github.com/ythosa/rating-list-monitoring-platform-api/pkg/apierrors"
	"github.com/ythosa/rating-list-monitoring-platform-api/pkg/logging"
	"github.com/ythosa/rating-list-monitoring-platform-api/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/service"
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
// @router /university/ [get].
func (u *UniversityImpl) GetAll(c *gin.Context) {
	universities, err := u.universityService.GetAll()
	if err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.NewAPIError(err))

		return
	}

	c.JSON(http.StatusOK, universities)
}

// Get
// @tags university
// @summary returns university by id
// @accept json
// @produce json
// @security AccessTokenHeader
// @param id path int true "university id"
// @success 200 {object} models.University
// @failure 400 {object} apierrors.APIError
// @failure 401 {object} apierrors.APIError
// @router /university/{id} [get].
func (u *UniversityImpl) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.InvalidQueryIDParam)

		return
	}

	university, err := u.universityService.GetByID(uint(id))
	if err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.NewAPIError(err))

		return
	}

	c.JSON(http.StatusOK, university)
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
