package db

import (
	"context"
	"fmt"
	"log"
	"time"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	MongoDB := os.Getenv("MONGO_URI")
	fmt.Println("connecting to MongoDB at:", MongoDB)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err !=nil {
		log.Fatal(err)
	}
	return client
	
}

var Client *mongo.Client = DBinstance()

func OpenCollection (client *mongo.Client, collecttionName string) *mongo.Collection {
	
	var collection *mongo.Collection = client.Database("real_chat").Collection(collecttionName)
	fmt.Println("collection opened")
	
	return collection
}
