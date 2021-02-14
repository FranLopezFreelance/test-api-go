package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TweetsResponse representa la estructura de datos de los tweets
type TweetsResponse struct {
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID string `bson:"userId,omitempty" json:"userId"`
	Message string `bson:"message" json:"message,omitempty"`
	Date time.Time `bson:"date" json:"date,omitempty"`
}