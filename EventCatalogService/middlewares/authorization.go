package middlewares

import (
	"eventCatalogService/grpc"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

type SignedDetails struct {
	Email string
	jwt.RegisteredClaims
}

const (
	AuthorizationHeader     = "Authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "payload"
)

func Authorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		authToken := context.GetHeader(AuthorizationHeader)
		if authToken == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization is not provided"})
			return
		}

		fields := strings.Fields(authToken)
		if len(fields) < 2 {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			return
		}

		AuthorizationType := strings.ToLower(fields[0])
		if AuthorizationType != AuthorizationTypeBearer {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unsupported authorization header"})
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
