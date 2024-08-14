package repository

import (
	"context"
	domain "task_manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db         mongo.Database
	collection mongo.Collection
}

func NewUserRepository(db mongo.Database) domain.UserRepository {
	return &userRepository{
		db: db, collection: *db.Collection("users"),
	}
}

func (ur *userRepository) CreateUser(user domain.User) error {
	data := bson.D{
		{"email", user.Email},
		{"password", user.Password},
		{"name", user.Name},
		{"role", user.Role}}
	_, err := ur.collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}
func (ur *userRepository) GetUserByEmail(email string) (domain.User, error) {
	filter := bson.D{{"email", email}}
	user := ur.collection.FindOne(context.TODO(), filter)
	if user.Err() != nil {
		return domain.User{}, user.Err()
	}
	var u domain.User
	user.Decode(&u)
	return u, nil
}

func (ur *userRepository) UpdateUser(email string, update domain.User) error {
	filter := bson.D{{"email", email}}
	data:=bson.D{{"$set",bson.D{
		{"email",update.Email},
		{"password",update.Password},
		{"name",update.Name},
		{"role",update.Role},
	}}}
	_, err := ur.collection.UpdateOne(context.TODO(), filter, data)
	if err != nil {
		return err
	}
	return nil
}
