package middleware

import (
	"net/http"
	"strings"

	"github.com/Darari17/golang-e-commerce/model"
	"github.com/Darari17/golang-e-commerce/security"
	"github.com/gin-gonic/gin"
)

type IAuthMiddleware interface {
	RequiredToken(roles ...string) gin.HandlerFunc
}

type authMiddleware struct {
	jwtHandler security.IJWTHandler
}

func NewAuthMiddleware(jwtHandler security.IJWTHandler) IAuthMiddleware {
	return &authMiddleware{
		jwtHandler: jwtHandler,
	}
}

// RequiredToken implements IAuthMiddleware.
func (a *authMiddleware) RequiredToken(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   "Token is required",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := a.jwtHandler.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   "Invalid token",
			})
			return
		}

		c.Set("user", model.User{
			ID:   claims.ID,
			Role: claims.Role,
		})

		if len(roles) == 0 {
			c.Next()
			return
		}

		for _, role := range roles {
			if role == claims.Role {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Forbidden resource",
		})
	}
}
