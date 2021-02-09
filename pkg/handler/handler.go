package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yongDataScince/rest-api/pkg/service"
)

type Handlers struct {
	services *service.Service
}

func NewHandlers(services *service.Service) *Handlers {
		return &Handlers{
			services: services,
		}
}

// InitRoutes ...
func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.upadateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItem)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.upadateItm)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}

	return router
}
