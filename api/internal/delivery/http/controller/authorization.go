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

type AuthorizationImpl struct {
	validate             *validator.Validate
	authorizationService service.Authorization
	logger               *logging.Logger
}

func NewAuthorizationImpl(validate *validator.Validate, authorizationService service.Authorization) *AuthorizationImpl {
	return &AuthorizationImpl{
		validate:             validate,
		authorizationService: authorizationService,
		logger:               logging.NewLogger("authorization controller"),
	}
}

// SignUp
// @tags authorization
// @summary sign up new user
// @description receives user credentials, creates user and returns user id
// @accept json
// @produce json
// @param payload body dto.SigningUp true "user credentials"
// @success 201 {object} dto.IDResponse
// @failure 400 {object} apierrors.APIError
// @failure 409 {object} apierrors.APIError
// @router /auth/sign-up [post].
func (a *AuthorizationImpl) SignUp(c *gin.Context) {
	var payload dto.SigningUp

	if err := c.BindJSON(&payload); err != nil {
		a.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.NewAPIError(err))

		return
	}

	if err := payload.Validate(a.validate); err != nil {
		a.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.NewAPIError(err))

		return
	}

	id, err := a.authorizationService.SignUpUser(payload)
	if err != nil {
		a.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusConflict, apierrors.NewAPIError(err))

		return
	}

	c.JSON(http.StatusCreated, dto.IDResponse{ID: id})
}

// SignIn
// @tags authorization
// @summary sign in user with jwt tokens response
// @description receives user credentials and returns jwt access and refresh tokens
// @accept json
// @produce json
// @param payload body dto.UserCredentials true "user credentials"
// @success 200 {object} dto.AuthorizationTokens
// @failure 400 {object} apierrors.APIError
// @failure 401 {object} apierrors.APIError
// @router /auth/sign-in [post].
func (a *AuthorizationImpl) SignIn(c *gin.Context) {
	var payload dto.UserCredentials

	if err := c.BindJSON(&payload); err != nil {
		a.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.NewAPIError(err))

		return
	}

	if err := payload.Validate(a.validate); err != nil {
		a.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.NewAPIError(err))

		return
	}

	tokens, err := a.authorizationService.GenerateTokens(payload)
	if err != nil {
		a.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, apierrors.NewAPIError(err))

		return
	}

	c.JSON(http.StatusOK, tokens)
}

// RefreshTokens
// @tags authorization
// @summary update jwt access and refresh tokens
// @description receives refresh token header and returns updated jwt access and refresh tokens
// @produce json
// @param RefreshToken header string true "refresh token header"
// @success 200 {object} dto.AuthorizationTokens
// @failure 401 {object} apierrors.APIError
// @router /auth/refresh-tokens [get].
func (a *AuthorizationImpl) RefreshTokens(c *gin.Context) {
	refreshToken, err := middleware.GetRefreshTokenFromRequest(c)
	if err != nil {
		a.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, apierrors.InvalidRefreshToken)

		return
	}

	tokens, err := a.authorizationService.RefreshTokens(refreshToken)
	if err != nil {
		a.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, apierrors.NewAPIError(err))

		return
	}

	c.JSON(http.StatusOK, tokens)
}

// Logout
// @tags authorization
// @summary logout user
// @description receives access token header and logouts user
// @accept json
// @produce json
// @security AccessTokenHeader
// @success 200 "logout success"
// @failure 400 {object} apierrors.APIError
// @failure 401 {object} apierrors.APIError
// @router /auth/logout [get].
func (a *AuthorizationImpl) Logout(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		a.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, apierrors.NewAPIError(err))

		return
	}

	accessToken, err := middleware.GetAccessTokenFromRequest(c)
	if err != nil {
		a.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, apierrors.InvalidAuthorizationHeader)

		return
	}

	if err := a.authorizationService.LogoutUser(userID, accessToken); err != nil {
		a.logger.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, apierrors.NewAPIError(err))

		return
	}

	c.Status(http.StatusOK)
}
