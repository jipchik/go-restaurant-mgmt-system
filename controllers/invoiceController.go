package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting invoices...")
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting invoice...")
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Creating invoice...")
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Updating invoice...")
	}
}
