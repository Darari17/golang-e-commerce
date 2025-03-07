package security

import (
	"errors"
	"time"

	"github.com/Darari17/golang-e-commerce/model"
	"github.com/golang-jwt/jwt/v5"
)

type jwtClaims struct {
	ID   uint
	Role string
	jwt.RegisteredClaims
}

type jwtHandler struct {
	secretKey []byte
}

type IJWTHandler interface {
	CreateToken(user *model.User) (string, error)
	VerifyToken(tokenString string) (*jwtClaims, error)
}

func NewJwtHandler(secretKey string) IJWTHandler {
	return &jwtHandler{
		secretKey: []byte(secretKey),
	}
}

// CreateToken implements IJWTHandler.
func (j *jwtHandler) CreateToken(user *model.User) (string, error) {
	claims := jwtClaims{
		ID:   user.ID,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Farid Rhamadhan Darari",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenString.SignedString(j.secretKey)
}

// VerifyToken implements IJWTHandler.
func (j *jwtHandler) VerifyToken(tokenString string) (*jwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}
