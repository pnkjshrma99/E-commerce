// controllers/product.go
package controllers

import (
	"net/http"

	"example.com/models"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	db models.DBHandler
}

func NewProductController(db models.DBHandler) *ProductController {
	return &ProductController{db: db}
}

func (pc *ProductController) ListProducts(c *gin.Context) {
	products, err := pc.db.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.db.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// Add more controller methods for updating and deleting products...
