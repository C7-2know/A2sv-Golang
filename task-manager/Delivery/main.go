package main

import (
	"context"
	"log"
	"os"
	router "task_manager/Delivery/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	route := gin.Default()
	er := godotenv.Load()
	if er != nil {
		log.Fatalf("Error loading .env file")
	}
	mongoURI := os.Getenv("MongoURI")
	if mongoURI == "" {
		log.Fatalf("MongoURI not found")
	}
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	db := client.Database("task_manager")
	router.SetUp(*db, route)
	route.Run(":8080")
}
