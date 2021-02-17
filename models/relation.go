package models

//Relation es la estructura de datos de relaciones entre usuarios
type Relation struct {
	FollowerID string `bson:"followerId" json:"followerId"`
	FollowedID string `bson:"followedId" json:"followedId"`
}