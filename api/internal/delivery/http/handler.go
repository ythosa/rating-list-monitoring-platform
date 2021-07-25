package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/ythosa/rating-list-monitoring-platfrom-api/docs"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/service"
)

type Handler struct {
	services *service.Service
	validate *validator.Validate
}

func NewHandler(services *service.Service, validate *validator.Validate) *Handler {
	return &Handler{
		services: services,
		validate: validate,
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
	}

	return router
}
