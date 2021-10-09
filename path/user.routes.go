package routes

import (
	"Instagram_Backend_API/controllers"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(incomingRoute *gin.Engine) {
	incomingRoute.POST("/users/", controllers.AddUser())
	incomingRoute.GET("/users/:user_id", controllers.GetUserById())

	// can be enabled to fetch auth tokens using the middleware.
	incomingRoute.POST("/users/login", controllers.LoginUser())
}
