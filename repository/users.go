package repository

import (
	"context"
	"fmt"
	"log"
)

type User struct {
	ID    string `bson:"_id,omitempty"`
	Name  string `bson:"name,omitempty"`
	Email string `bson:"email,omitempty"`
	Age   int    `bson:"age,omitempty"`
}

type Users []User

func (db *MongoDatabase) CreateUser(user User) error {
	ctx := context.Background()
	_, err := db.Db.Collection("users").InsertOne(ctx, &user)
	if err != nil {
		return err
	}
	return nil
}

func (db *MongoDatabase) GetUsersByIds(ctx context.Context, ids []string) (Users, error) {
	var users Users
	fmt.Println(ids)
	err := db.FindByIds(ctx, "users", ids, &users)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return users, nil
}
