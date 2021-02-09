package models

import "time"

// Tweet representa la estructura de datos de un tweet
type Tweet struct {
	UserID string `bson:"_id,omitempty" json:"id"`
	Message string `bson:"message" json:"message,omitempty"`
	Date time.Time `bson:"date" json:"date,omitempty"`
}