package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/service"
)

type Authorization interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	RefreshTokens(c *gin.Context)
	Logout(c *gin.Context)
}

type Controller struct {
	Authorization
}

func NewController(validate *validator.Validate, services *service.Service) *Controller {
	return &Controller{
		Authorization: NewAuthorizationImpl(validate, services.Authorization),
	}
}