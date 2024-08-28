package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Type     string             `json:"type"`
	Age      int                `json:"age"`
	IsActive bool               `json:"is_active"`
}
