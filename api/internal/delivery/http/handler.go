package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ythosa/rating-list-monitoring-platform-api/docs" // swagger documentation
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/delivery/http/controllers"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/services"
	"github.com/ythosa/rating-list-monitoring-platform-api/pkg/middleware"
)

type Handler struct {
	services    *services.Service
	validate    *validator.Validate
	controllers *controllers.Controller
}

func NewHandler(services *services.Service, validate *validator.Validate) *Handler {
	return &Handler{
		services:    services,
		validate:    validate,
		controllers: controllers.NewController(validate, services),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(getCORSConfig())
	router.Use(middleware.NewMetricsMiddleware("/metrics").Metrics())

	router.GET("/metrics", prometheusHandler())

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
			university.GET("/", h.controllers.University.GetAll)
			university.GET("/:id", h.controllers.University.Get)
			university.GET("/get_for_user", h.controllers.University.GetForUser)
			university.POST("/set_for_user", h.controllers.University.SetForUser)
		}

		direction := api.Group("/direction", middleware.UserIdentity)
		{
			direction.GET("/", h.controllers.Direction.GetAll)
			direction.GET("/:id", h.controllers.Direction.Get)
			direction.GET("/get_for_user", h.controllers.Direction.GetForUser)
			direction.GET("/get_for_user_with_rating", h.controllers.Direction.GetForUserWithRating)
			direction.POST("/set_for_user", h.controllers.Direction.SetForUser)
		}
	}

	return router
}

func getCORSConfig() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{
		"Access-Control-Allow-Headers",
		"Origin",
		"Accept",
		"X-Requested-With",
		"Content-Type",
		"Access-Control-Request-Method",
		"Access-Control-Request-Headers",
		"AuthTokens",
	}

	return cors.New(config)
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
