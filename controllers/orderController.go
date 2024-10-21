package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting orders...")
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting order...")
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Creating order...")
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Updating order...")
	}
}
