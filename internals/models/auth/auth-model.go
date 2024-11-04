package auth

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserInfoSignUp struct {
	ID       primitive.ObjectID `json:"id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

type UserRepository interface {
	AddUser(user *UserInfoSignUp) error
}
type MongoUserRepository struct {
	Collection *mongo.Collection
}

func (repo *MongoUserRepository) AddUser(user *UserInfoSignUp) error {
	user.ID = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.Collection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("failed to insert user: %v", err)
		return err
	}
	log.Println("user created successfully")
	return nil
}
