package db

import (
	"context"
	"fmt"
	"time"

	"github.com/FranLopezFreelance/models"
	"go.mongodb.org/mongo-driver/bson"
)

//GetRelation se encarga de checkear si sigo o no a un usuario
func GetRelation(t models.Relation) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoClient.Database("test-api-go")
	col := db.Collection("relations")

	condition := bson.M{
		"followerId": t.FollowerID,
		"followedId": t.FollowedID,
	}

	var result models.Relation
	fmt.Println(result)
	err := col.FindOne(ctx, condition).Decode(&result)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return true, nil

}