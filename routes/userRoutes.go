package routes

import (
	"github.com/gin-gonic/gin"

	controller "restaurant-management-system/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:userId", controller.GetUser())
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}
