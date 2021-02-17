package db

import (
	"context"
	"time"

	"github.com/FranLopezFreelance/models"
)

//CreateRelation se encarga de crear una relación entre usuarios
func CreateRelation(t models.Relation) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoClient.Database("test-api-go")
	col := db.Collection("relations")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}