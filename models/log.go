package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Device  string             `json:"name" bson:"device"`
	LogedAt time.Time          `json:"publishedAt" bson:"logedAt"`
}
