package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"restaurant-management-system/database"
	"restaurant-management-system/middleware"
	"restaurant-management-system/routes"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.NoteRoutes(router)
	routes.OrderItemRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)
	routes.UserRoutes(router)

	// Listen and serve on 0.0.0.0:8080
	router.Run(":" + port)
}
