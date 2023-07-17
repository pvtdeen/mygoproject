package api

import (
	"mygoproject/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type server struct {
	db *repository.MongoDatabase
}

func InitializeServer(conn *mongo.Client) server {
	xdb := repository.NewDB("MyDatabase", conn)
	return server{
		db: xdb,
	}
}
