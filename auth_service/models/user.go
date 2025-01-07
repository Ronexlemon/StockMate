package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json :"id"`
	Username  string             `bson:"username" json:"username"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Role      string             `bson:"role"json:"role"`
	CreatedAt string             `bson:"created_at" json:"created_at"`
	UpdatedAt string             `bson:"updated_at" json:"updated_at"`
}
