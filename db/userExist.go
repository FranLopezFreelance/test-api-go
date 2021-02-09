package db

import (
	"context"
	"time"

	"github.com/FranLopezFreelance/models"
	"go.mongodb.org/mongo-driver/bson"
)

// UserExist checkea si el usuario existe
func UserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	
	defer cancel()
	
	db := MongoClient.Database("test-api-go")
	col := db.Collection("users")

	condition := bson.M{"email":email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
