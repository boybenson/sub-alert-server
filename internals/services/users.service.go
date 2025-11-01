package services

import (
	"context"
	"example/sub-alert-server/internals/database"
	"example/sub-alert-server/internals/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers() []models.User {
	var users []models.User

	cursor, err := database.UserCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error fetching users:", err)
		return users
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Println("Error decoding user:", err)
			continue
		}
		users = append(users, user)
	}

	return users
}