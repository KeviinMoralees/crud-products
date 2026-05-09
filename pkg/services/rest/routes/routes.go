package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/KeviinMoralees/crud-products/pkg/domains/product"
	"github.com/KeviinMoralees/crud-products/pkg/services/rest/handlers"
)

type Handler struct {
	router  *gin.Engine
	service product.Service
}

func NewHandler(service product.Service) *Handler {
	h := &Handler{
		router:  gin.Default(),
		service: service,
	}
	h.registerRoutes()
	return h
}

func (h *Handler) Run(addr string) {
	h.router.Run(addr)
}

func (h *Handler) registerRoutes() {
	api := h.router.Group("/api/productos")
	api.POST("", handlers.CreateProduct(h.service))
	api.GET("", handlers.GetAllProducts(h.service))
	api.GET("/:id", handlers.GetProductByID(h.service))
	api.PUT("/:id", handlers.UpdateProduct(h.service))
	api.DELETE("/:id", handlers.DeleteProduct(h.service))
}
