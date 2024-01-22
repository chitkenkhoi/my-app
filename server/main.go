package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"server/middleware"
	Mongo "server/mongo"
	"server/routes"
)

func main() {
	Mongo.ConnectDB()
	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
	}
	alloworigins := os.Getenv("ALLOWORIGINS")
	fmt.Print(alloworigins[0])
	defer func() {
		if err := Mongo.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	////connect to firebase database
	router := gin.Default()
	////// gin framework
	middleware.Cors(router)
	////// CORS
	routes.Routes(router)
	//////routes for API
	router.Run(":8080")
}
