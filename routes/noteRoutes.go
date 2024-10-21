package routes

import (
	"github.com/gin-gonic/gin"

	controller "restaurant-management-system/controllers"
)

func NoteRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/notes", controller.GetNotes())
	incomingRoutes.GET("/notes/:noteId", controller.GetNote())
	incomingRoutes.POST("/notes", controller.CreateNote())
	incomingRoutes.PATCH("/notes/:noteId", controller.UpdateNote())
}
