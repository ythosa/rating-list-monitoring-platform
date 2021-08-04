package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/service"
)

type Authorization interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	RefreshTokens(c *gin.Context)
	Logout(c *gin.Context)
}

type User interface {
	GetUsername(c *gin.Context)
	GetProfile(c *gin.Context)
}

type University interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	GetForUser(c *gin.Context)
	SetForUser(c *gin.Context)
}

type Direction interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	GetForUser(c *gin.Context)
	GetForUserWithRating(c *gin.Context)
	SetForUser(c *gin.Context)
}

type Controller struct {
	Authorization
	User
	University
	Direction
}

func NewController(validate *validator.Validate, services *service.Service) *Controller {
	return &Controller{
		Authorization: NewAuthorizationImpl(validate, services.Authorization),
		User:          NewUserImpl(validate, services.User),
		University:    NewUniversityImpl(validate, services.University),
		Direction:     NewDirectionImpl(validate, services.Direction),
	}
}
