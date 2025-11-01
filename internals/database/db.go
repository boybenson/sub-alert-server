package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	mongoURI := os.Getenv("MONGO_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("MongoDB client creation error:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	DB = client.Database("SubAlert")
	log.Println("✅ Connected to MongoDB!")
}


var (
	UserCollection         *mongo.Collection
	SubscriptionCollection *mongo.Collection
)

func InitCollections() {
	UserCollection = DB.Collection("users")
	SubscriptionCollection = DB.Collection("subscriptions")
}