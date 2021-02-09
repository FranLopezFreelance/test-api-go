package db

import (
	"context"
	"fmt"
	"time"

	"github.com/FranLopezFreelance/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SearchProfile busca y devuelve un perfil de usuario
func SearchProfile(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoClient.Database("test-api-go")
	col := db.Collection("users")

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("registro no encontrado "+err.Error())
		return profile, err
	}

	return profile, nil
}
