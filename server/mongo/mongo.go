package mongo

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var Client *mongo.Client

func ConnectDB() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	if err := godotenv.Load(); err != nil {
		fmt.Println("Failed to load from environment")
	}
	username := os.Getenv("USERNAME_MONGO")
	password := os.Getenv("PASSWORD")
	uri := "mongodb+srv://" + username + ":" + password + "@cluster0.buwnaid.mongodb.net/?retryWrites=true&w=majority"
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	Client = client
}
