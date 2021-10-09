package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID         primitive.ObjectID `bson:"_id"`
	Caption    string             `json:"caption" validate:"required,max=255"`
	ImageURL   string             `json:"image_url" validate:"required"`
	Post_id    string             `json:"name"`
	User_id    string             `json:"user_id"`
	Created_at time.Time          `json:"created_at"`
}
