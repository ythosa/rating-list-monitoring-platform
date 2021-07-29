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

type DirectionImpl struct {
	validate         *validator.Validate
	directionService service.Direction
	logger           *logging.Logger
}

func NewDirectionImpl(validate *validator.Validate, directionService service.Direction) *DirectionImpl {
	return &DirectionImpl{
		validate:         validate,
		directionService: directionService,
		logger:           logging.NewLogger("direction controller"),
	}
}

// GetAll
// @tags direction
// @summary returns all directions
// @accept json
// @produce json
// @security AccessTokenHeader
// @success 200 {object} map[string][]dto.Direction
// @failure 400 {object} apierrors.APIError
// @failure 401 {object} apierrors.APIError
// @router /direction/get_all [get].
func (u *DirectionImpl) GetAll(c *gin.Context) {
	_, err := middleware.GetUserID(c)
	if err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, apierrors.NewAPIError(err))

		return
	}

	directions, err := u.directionService.GetAll()
	if err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.NewAPIError(err))

		return
	}

	c.JSON(http.StatusOK, directions)
}

// Get
// @tags direction
// @summary returns user directions
// @accept json
// @produce json
// @security AccessTokenHeader
// @success 200 {object} map[string][]dto.Direction
// @failure 400 {object} apierrors.APIError
// @failure 401 {object} apierrors.APIError
// @router /direction/get [get].
func (u *DirectionImpl) Get(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, apierrors.NewAPIError(err))

		return
	}

	directions, err := u.directionService.Get(userID)
	if err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.NewAPIError(err))

		return
	}

	c.JSON(http.StatusOK, directions)
}

// Set
// @tags direction
// @summary set directions to user
// @description receives direction ids and sets it to user
// @accept json
// @produce json
// @security AccessTokenHeader
// @param payload body dto.IDs true "direction ids"
// @success 200 "success"
// @failure 400 {object} apierrors.APIError
// @failure 401 {object} apierrors.APIError
// @router /direction/set [post].
func (u *DirectionImpl) Set(c *gin.Context) {
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

	if err := u.directionService.Set(userID, payload); err != nil {
		u.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.NewAPIError(err))

		return
	}

	c.Status(http.StatusOK)
}
