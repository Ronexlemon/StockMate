package service

import (
	"context"
	"microservice/auth_service/database"
	"microservice/auth_service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func Register(user models.User,ctx context.Context)(*mongo.InsertOneResult,error){
	// register service
	user.CreatedAt = time.Now().Format(time.RFC3339)
	user.UpdatedAt = time.Now().Format(time.RFC3339)
	return database.UserCollection.InsertOne(ctx,user)
	
}

