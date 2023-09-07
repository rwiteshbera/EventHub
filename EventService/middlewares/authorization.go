package middlewares

import (
	"eventCatalogService/grpc"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type SignedDetails struct {
	Email string
	jwt.RegisteredClaims
}

const (
	AuthorizationCookieName = "Authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "payload"
)

func Authorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		authToken, err := context.Cookie(AuthorizationCookieName)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization is not provided"})
			return
		}

		fields := strings.Fields(authToken)
		if len(fields) < 2 {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization"})
			return
		}

		AuthorizationType := strings.ToLower(fields[0])
		if AuthorizationType != AuthorizationTypeBearer {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unsupported authorization"})
			return
		}

		AccessToken := fields[1]
		payload, err := grpc.GRPCServe(AccessToken)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		context.Set(AuthorizationPayloadKey, payload.UserEmail)
		context.Next()
	}
}
