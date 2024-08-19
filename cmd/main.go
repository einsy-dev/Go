package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"app/internal/movies"
	"app/internal/users"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL")) // MongoDB client options
	client, err := mongo.Connect(context.TODO(), clientOptions)        // Connect to MongoDB
	client.Ping(context.TODO(), nil)                                   // Ping MongoDB
	db := client.Database(os.Getenv("DATABASE"))                       // Select database

	if err != nil {
		log.Fatal("Error connecting to MongoDB", err)
	}

	app := gin.Default()
	app.Static("/static", "./public")

	movies.Router(app, db.Collection("movies"))
	users.Router(app, db.Collection("users"))

	app.Run()
}
