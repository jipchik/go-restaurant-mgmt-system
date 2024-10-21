package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting foods...")
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting food...")
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Creating food...")
	}
}

func round(num float64) int {
	return 1
}

func toFixed(num float64, precision int) float64 {
	return 1.22
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Updating food...")
	}
}
