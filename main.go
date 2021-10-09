package main

import (
	"os"

	"Instagram_Backend_API/path"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	path.AddUserRoutes(router)
	path.AddPostRoutes(router)

	router.Run(":" + port)
}
