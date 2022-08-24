package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Upload struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	ImageName string             `json:"name" bson:"image_name"`
}
