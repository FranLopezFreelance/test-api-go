package db

import (
	"context"
	"time"

	"github.com/FranLopezFreelance/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdateProfile modifica datos de perfil de un usuario
func UpdateProfile(user models.User, id string) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoClient.Database("test-api-go")
	col := db.Collection("users")

	newData := make(map[string] interface{})

	newData["birthDate"] = user.BirthDate

	if len(user.Name) > 0 {
		newData["name"] = user.Name
	}

	if len(user.LastName) > 0 {
		newData["lastName"] = user.LastName
	}

	if len(user.Biography) > 0 {
		newData["biography"] = user.Biography
	}

	if len(user.Location) > 0 {
		newData["location"] = user.Location
	}

	if len(user.WebSite) > 0 {
		newData["avatar"] = user.WebSite
	}

	if len(user.Avatar) > 0 {
		newData["avatar"] = user.Avatar
	}

	if len(user.Banner) > 0 {
		newData["avatar"] = user.Banner
	}

	updateData := bson.M{
		"$set": newData,
	}

	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": bson.M{ "$eq": objectID }}

	_, err := col.UpdateOne(ctx, filter, updateData)

	if err != nil {
		return false, err
	}

	return true, nil
}
