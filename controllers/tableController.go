package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting tables...")
	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting table...")
	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Creating table...")
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Updating table...")
	}
}
