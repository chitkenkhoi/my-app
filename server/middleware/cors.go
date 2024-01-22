package middleware

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func Cors(router *gin.Engine) {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
	}
	AllowOrigins := strings.Split(os.Getenv("ALLOWORIGINS"),",")
	router.Use(cors.New(cors.Config{
		AllowOrigins:     AllowOrigins,
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Access-Control-Allow-Credentials"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
	}))
}
