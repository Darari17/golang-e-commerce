package handler

import (
	"net/http"

	"github.com/Darari17/golang-e-commerce/dto"
	"github.com/Darari17/golang-e-commerce/model"
	"github.com/Darari17/golang-e-commerce/service"
	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	Register(c *gin.Context)
	Profile(c *gin.Context)
}

type userHandler struct {
	userService service.IUserService
}

func NewUserHandler(userService service.IUserService) IUserHandler {
	return &userHandler{
		userService: userService,
	}
}

// Profile implements IUserHandler.
func (u *userHandler) Profile(c *gin.Context) {
	value, exists := c.Get("user")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, ok := value.(model.User)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}

	profile, err := u.userService.Profile(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   profile,
	})
}

// Register implements IUserHandler.
func (u *userHandler) Register(c *gin.Context) {
	var payload dto.RegisterRequest
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	register, err := u.userService.Register(payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Created",
		"data":   register,
	})
}
