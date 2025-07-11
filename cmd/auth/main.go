package main

import (
	"fmt"
	"go-gin-api/internal/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	authService := auth.NewAuthService()
	r.POST("/auth/login", authService.Login)

	fmt.Println("Auth Service is running at :8081")
	r.Run(":8081") // Auth Service on port 8081

	r.POST("/auth/verify", authService.VerifyToken)
}
