package db

import (
	"context"
	"log"
	"time"

	"github.com/FranLopezFreelance/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetTweetsByUser se necarga de mostrar tweetes de un usuario
func GetTweetsByUser(id string, page int64) ([]*models.TweetsResponse, bool){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoClient.Database("test-api-go")
	col := db.Collection("tweets")

	var tweets []*models.TweetsResponse

	condition := bson.M{
		"userId": id,
	}

	var limit int64 = 20

	options := options.Find()
	options.SetLimit(limit)
	options.SetSort(bson.D{{Key:"date", Value: -1}})
	options.SetSkip((page - 1) * 20)

	collection, err := col.Find(ctx, condition, options)

	if err != nil {
		log.Fatal(err.Error())
		return tweets, false
	}

	for collection.Next(context.TODO()){
		var item models.TweetsResponse
		err := collection.Decode(&item)
		if err != nil {
			return tweets, false
		}
		tweets = append(tweets, &item)
	}

	return tweets, true
}
