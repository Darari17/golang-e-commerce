package handler

import (
	"net/http"

	"github.com/Darari17/golang-e-commerce/dto"
	"github.com/Darari17/golang-e-commerce/service"
	"github.com/gin-gonic/gin"
)

type IAuthHandler interface {
	Login(c *gin.Context)
}

type authHandler struct {
	authService service.IAuthService
}

func NewAuthService(authService service.IAuthService) IAuthHandler {
	return &authHandler{
		authService: authService,
	}
}

// Login implements IAuthHandler.
func (a *authHandler) Login(c *gin.Context) {
	var payload dto.LoginRequest
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := a.authService.Login(payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Created",
		"token":  token,
	})
}
