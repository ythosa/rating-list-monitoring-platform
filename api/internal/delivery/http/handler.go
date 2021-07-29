package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/ythosa/rating-list-monitoring-platfrom-api/docs"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/delivery/http/controller"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/delivery/http/middleware"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/service"
)

type Handler struct {
	services    *service.Service
	validate    *validator.Validate
	controllers *controller.Controller
}

func NewHandler(services *service.Service, validate *validator.Validate) *Handler {
	return &Handler{
		services:    services,
		validate:    validate,
		controllers: controller.NewController(validate, services),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		swaggerDocumentation := api.Group("/docs")
		{
			swaggerDocumentation.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		authorization := api.Group("/auth")
		{
			authorization.POST("/sign-up", h.controllers.Authorization.SignUp)
			authorization.POST("/sign-in", h.controllers.Authorization.SignIn)
			authorization.GET("/refresh-tokens", h.controllers.Authorization.RefreshTokens)
			authorization.GET("/logout", middleware.UserIdentity, h.controllers.Authorization.Logout)
		}

		user := api.Group("/user", middleware.UserIdentity)
		{
			user.GET("/get_username", h.controllers.User.GetUsername)
			user.GET("/get_profile", h.controllers.User.GetProfile)
		}

		university := api.Group("/university", middleware.UserIdentity)
		{
			university.GET("/get_all", h.controllers.University.GetAll)
			university.GET("/get", h.controllers.University.Get)
			university.POST("/set", h.controllers.University.Set)
		}

		direction := api.Group("/direction", middleware.UserIdentity)
		{
			direction.GET("/get_all", h.controllers.Direction.GetAll)
			direction.GET("/get", h.controllers.Direction.Get)
			direction.POST("/set", h.controllers.Direction.Set)
		}
	}

	return router
}
