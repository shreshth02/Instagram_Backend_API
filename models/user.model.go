package models

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `json:"name" validate:"required,min=3,max=255"`
	Email    string             `json:"email" validate:"email,required"`
	Password string             `json:"password" validate:"required,min=6"`
	User_id  string             `json:"user_id"`
}

type LoginBody struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=6"`
}

type SignedClaims struct {
	UserId string
	jwt.StandardClaims
}

func (user *User) GenerateJWT() (string, error) {
	expirationTime := time.Now().Add(15 * time.Hour)
	claims := &SignedClaims{
		UserId: user.User_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		log.Panic(err)
		return "Could not generated user token.", nil
	}
	return token, err
}
