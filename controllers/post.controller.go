package controllers

import (
	"context"
	"net/http"
	"time"

	"Instagram_Backend_API/database"
	"Instagram_Backend_API/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var postCollection *mongo.Collection = database.OpenCollection(database.Client, "post")

func AddPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var post models.Post
		if err := c.BindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		err := validate.Struct(post)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		post.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		post.ID = primitive.NewObjectID()
		post.Post_id = post.ID.Hex()
		post.User_id = c.GetString("user_id")

		defer cancel()
		resultInsertionNumber, insertErr := postCollection.InsertOne(ctx, post)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user."})
			return
		}
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

func GetPostById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		PostId := c.Param("post_id")
		var post models.Post
		if err := postCollection.FindOne(ctx, bson.M{"post_id": PostId}).Decode(&post); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid id/User not found."})
			return
		}

		c.JSON(http.StatusOK, post)
	}
}

func GetPostsOfUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		UserId := c.Param("user_id")

		filter := bson.M{"user_id": UserId}

		cursor, err := postCollection.Find(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong."})
			return
		}
		var posts []models.Post
		if err = cursor.All(ctx, &posts); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong."})
			return
		}
		c.JSON(http.StatusOK, posts)
	}
}
