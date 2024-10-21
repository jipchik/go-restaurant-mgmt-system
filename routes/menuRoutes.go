package routes

import (
	"github.com/gin-gonic/gin"

	controller "restaurant-management-system/controllers"

)

func MenuRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/menus", controller.GetMenus())
	incomingRoutes.GET("/menus/:menuId", controller.GetMenu())
	incomingRoutes.POST("/menus", controller.CreateMenu())
	incomingRoutes.PATCH("/menus/:menuId", controller.UpdateMenu())
}
