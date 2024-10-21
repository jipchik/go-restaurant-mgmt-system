package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting menus...")
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting menu...")
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Creating menu...")
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Updating menu...")
	}
}
