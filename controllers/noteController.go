package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting notes...")
	}
}

func GetNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting note...")
	}
}

func CreateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Creating note...")
	}
}

func UpdateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Updating note...")
	}
}
