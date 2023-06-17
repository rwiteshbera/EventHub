package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type SignedDetails struct {
	Email string
	jwt.RegisteredClaims
}

func GenerateToken(email string, jwtSecret string) (signedToken string, err error) {
	claims := SignedDetails{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 168)), // Token will be expired after 168 H
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jwtSecret))
	if err != nil {
		fmt.Println("error ", err)
		return
	}

	return token, err
}

// Validate JWT Token
func validateToken(signedToken string, jwtSecret string) (claims *SignedDetails, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*SignedDetails)

	if !ok {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func Authenticate(jwtSecret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("authorization")

		if clientToken == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no authorization header provided"})
			ctx.Abort()
			return
		}

		claims, err := validateToken(clientToken, jwtSecret)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ctx.Abort()
		}

		ctx.Set("email", claims.Email)
		ctx.Set("uid", claims.ID)
		ctx.Next()
	}
}
