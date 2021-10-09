package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"Instagram_Backend_API/database"
	"Instagram_Backend_API/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func ValidateToken(signedToken string) (claims *models.SignedClaims, msg string) {
	msg = ""
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.SignedClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*models.SignedClaims)
	if !ok {
		msg = fmt.Sprintf("the token is invalid.")
		msg = err.Error()
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired.")
		msg = err.Error()
		return
	}
	return claims, msg
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Access denied. No token provided.")})
			c.Abort()
			return
		}

		claims, err := ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		var foundUser models.User
		userError := userCollection.FindOne(ctx, bson.M{"user_id": claims.UserId}).Decode(&foundUser)
		defer cancel()
		if userError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Access denied. Token Invalid."})
			c.Abort()
			return
		}
		c.Set("user_id", foundUser.User_id)
		c.Next()
	}
}
