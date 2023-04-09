package controllers

import (
	"context"
	"errors"
	"net/http"
	"time"
	"userService/api"
	"userService/database"
	"userService/models"
	"userService/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var tokenExpiry time.Duration = 5 * time.Minute // Token Expiry
var sessionIdMaxAge int = 2 * 60                // 2 minutes

func Signup(server *api.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userLoginRequest models.UserLogin

		err1 := ctx.ShouldBindJSON(&userLoginRequest)
		if err1 != nil {
			LogError(ctx, http.StatusBadRequest, err1)
			return
		}

		if userLoginRequest.Name == "" || userLoginRequest.Email == "" {
			LogError(ctx, http.StatusBadRequest, errors.New("please make sure all fields are filled in correctly"))
			return
		}

		redisDB := database.CreateRedisClient(&server.Config, 0)
		defer redisDB.Close()

		// Generate OTP
		otp, err2 := utils.GenerateOTP()
		if err2 != nil {
			LogError(ctx, http.StatusInternalServerError, err2)
			return
		}

		err3 := redisDB.Set(database.Ctx, userLoginRequest.Email, otp, tokenExpiry).Err()
		if err3 != nil {
			LogError(ctx, http.StatusInternalServerError, err3)
			return
		}

		ctx.SetCookie("name", userLoginRequest.Name, 0, "/", server.Config.SERVER_HOST, false, true)
		ctx.SetCookie("email", userLoginRequest.Email, 0, "/", server.Config.SERVER_HOST, false, true)

		LogMessage(ctx, userLoginRequest)
	}
}

func VerifyOTP(server *api.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userOtp models.OTP

		err0 := ctx.ShouldBindJSON(&userOtp)
		if err0 != nil {
			LogError(ctx, http.StatusBadRequest, err0)
			return
		}

		currentUserName, err0 := ctx.Cookie("name")
		if err0 != nil {
			LogError(ctx, http.StatusBadRequest, err0)
			return
		}
		currentUserEmail, err0 := ctx.Cookie("email")
		if err0 != nil {
			LogError(ctx, http.StatusBadRequest, err0)
			return
		}

		dbContext, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()

		// Create MongoInstance
		mongoClient, err2 := database.CreateMongoInstance(server.Config.MONGO_DB_URI)
		if err2 != nil {
			LogError(ctx, http.StatusInternalServerError, err2)
			return
		}

		userCollection := database.OpenMongoCollection(mongoClient, "user")

		// Check the user exists or not
		filter := bson.D{{Key: "email", Value: currentUserEmail}}

		count, err := userCollection.CountDocuments(ctx, filter)
		if err != nil {
			LogError(ctx, http.StatusInternalServerError, err)
			return
		}

		// Insert user data in MongoDB if it is not present
		if count < 1 {
			_, err3 := userCollection.InsertOne(dbContext, models.UserLogin{Name: currentUserName, Email: currentUserEmail})
			if err3 != nil {
				LogError(ctx, http.StatusInternalServerError, err3)
				return
			}
		}

		// Create Client in Redis
		redisDB := database.CreateRedisClient(&server.Config, 0)
		defer redisDB.Close()

		// Check if email is present
		isEmailPresent, err4 := redisDB.Exists(ctx, currentUserEmail).Result()
		if err4 != nil {
			LogError(ctx, http.StatusInternalServerError, err4)
			return
		}

		if isEmailPresent == 1 {
			savedOTP, err5 := redisDB.Get(ctx, currentUserEmail).Result()
			if err5 != nil {
				LogError(ctx, http.StatusInternalServerError, errors.New("otp has expireds"))
				return
			}

			sessionId := utils.GenerateSessionId(currentUserEmail)
			if savedOTP == userOtp.Otp {
				redisDB.Del(ctx, currentUserEmail)
				ctx.JSON(http.StatusOK, gin.H{"message": "verified", "sessionId": sessionId, "count": count})

				// set a cookie for authorization
				ctx.SetCookie("session", sessionId, sessionIdMaxAge, "/", server.Config.SERVER_HOST, false, true)
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "wrong otp"})
			}
		}

		ctx.JSON(http.StatusOK, count)
	}
}
