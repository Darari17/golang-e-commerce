package handler

import (
	"net/http"
	"strconv"

	"github.com/Darari17/golang-e-commerce/dto"
	"github.com/Darari17/golang-e-commerce/service"
	"github.com/gin-gonic/gin"
)

type IProductHandler interface {
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	FindProductById(c *gin.Context)
	FindAllProducts(c *gin.Context)
}

type productHandler struct {
	productService service.IProductService
}

func NewProducyHandler(productService service.IProductService) IProductHandler {
	return &productHandler{
		productService: productService,
	}
}

// CreateProduct implements IProductHandler.
func (p *productHandler) CreateProduct(c *gin.Context) {
	var payload dto.ProductRequest
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createProduct, err := p.productService.CreateProduct(payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Created",
		"data":   createProduct,
	})
}

// DeleteProduct implements IProductHandler.
func (p *productHandler) DeleteProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = p.productService.DeleteProduct(uint(productId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Delete successfully",
	})
}

// FindAllProducts implements IProductHandler.
func (p *productHandler) FindAllProducts(c *gin.Context) {
	products, err := p.productService.FindAllProducts()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   products,
	})
}

// FindProductById implements IProductHandler.
func (p *productHandler) FindProductById(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := p.productService.FindProductById(uint(productId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   product,
	})
}

// UpdateProduct implements IProductHandler.
func (p *productHandler) UpdateProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var payload dto.ProductUpdateRequest
	err = c.ShouldBindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateProduct, err := p.productService.UpdateProduct(uint(productId), payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   updateProduct,
	})
}
