package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting order items...")
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting order item...")
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting order items by order...")
	}
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {
	itemsInOrder := []primitive.M
	errorReceived := error
	return itemsInOrder, errorReceived
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Creating order item...")
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Updating order item...")
	}
}
