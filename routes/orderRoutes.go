package routes

import (
	"github.com/gin-gonic/gin"

	controller "restaurant-management-system/controllers"

)

func OrderRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/orders/:orderId", controller.GetOrder())
	incomingRoutes.POST("/orders", controller.CreateOrder())
	incomingRoutes.PATCH("/orders/:orderId", controller.UpdateOrder())
}
