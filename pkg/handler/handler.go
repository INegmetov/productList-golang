package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/inegmetov/productList-golang/docs"
	"github.com/inegmetov/productList-golang/pkg/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		lists := api.Group("/list")
		{
			lists.POST("/", h.createList)
			lists.GET("/:id", h.getListById)

			items := lists.Group(":id/products")
			{
				items.GET("/", h.getAll)
				items.POST("/", h.createItem)
			}
		}
		items := api.Group("/products")
		{
			items.GET("/", h.getAllItems)
		}
	}
	return router

}
