package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


// var client *mongo.Client
var  UserCollection *mongo.Collection

func Loadenv()string{
	err := godotenv.Load(".env")
	if err !=nil{
		log.Fatal(err)
	}
	mongo_url := os.Getenv("MONGO_URL")
	return mongo_url
}

func Connect()*mongo.Client{
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(Loadenv())
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		}
		UserCollection =client.Database("users_service").Collection("user_auth")
		return client
		

}
// func GetCollection() *mongo.Collection {
// 	client := Connect()
// 	// Return the collection from the database
// 	collection := client.Database("users_service").Collection("user_auth")
// 	return collection
// }