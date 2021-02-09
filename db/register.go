package db

import (
	"context"
	"time"

	"github.com/FranLopezFreelance/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Register Registra un usuario
func Register(user models.User) (string, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	
	defer cancel()

	db := MongoClient.Database("test-api-go")
	col := db.Collection("users")

	user.Password, _ = PasswordEncrypt(user.Password)

	result, err := col.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
