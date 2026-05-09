package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KeviinMoralees/crud-products/pkg/domains/entities"
	"github.com/KeviinMoralees/crud-products/pkg/domains/product"
)

func CreateProduct(svc product.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p entities.Product
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		created, err := svc.CreateProduct(c.Request.Context(), p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, created)
	}
}
