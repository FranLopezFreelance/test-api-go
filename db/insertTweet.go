package db

import (
	"context"
	"time"

	"github.com/FranLopezFreelance/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertTweet agrega un tweet a la base
func InsertTweet(tweet models.Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoClient.Database("test-api-go")
	col := db.Collection("tweets")

	newTweet := bson.M{
		"userId": tweet.UserID,
		"message": tweet.Message,
		"date": tweet.Date,
	}

	result, err := col.InsertOne(ctx, newTweet)

	if err != nil {
		return string(""), false, err
	}

	objectID, _ := result.InsertedID.(primitive.ObjectID)

	return objectID.String(),true , nil
}
