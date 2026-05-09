package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KeviinMoralees/crud-products/pkg/domains/product"
)

func GetAllProducts(svc product.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := svc.GetAllProducts(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}
