package routes

import (
	"Instagram_Backend_API/controllers"
	"Instagram_Backend_API/middleware"

	"github.com/gin-gonic/gin"
)

func AddPostRoutes(incomingRoute *gin.Engine) {
	incomingRoute.Use(middleware.Auth())
	incomingRoute.POST("/posts/", controllers.AddPost())
	incomingRoute.GET("/posts/users/:user_id", controllers.GetPostsOfUser())
	incomingRoute.GET("/posts/:post_id", controllers.GetPostById())
}
