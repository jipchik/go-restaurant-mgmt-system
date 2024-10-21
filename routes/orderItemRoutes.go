package routes

import (
	"github.com/gin-gonic/gin"

	controller "restaurant-management-system/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItemId", controller.GetOrderItem())
	incomingRoutes.GET("/orderItems/order/:orderId", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItemId", controller.UpdateOrderItem())
}
