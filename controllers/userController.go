package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting users...")
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Getting user...")
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Logging user in...")
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Signing user up...")
	}
}

func HashPassword(password string) string {
	fmt.Println("Hashing password...")
	return ""
}

func VerifyPassword(usersPassword string, providedPassword string) (bool, string) {
	fmt.Println("Verifying password...")
	return true, ""
}
