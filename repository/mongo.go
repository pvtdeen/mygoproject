package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenConnection() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://gtmhubdbuser:233buE5xDmSiSrKp@127.0.0.1:27017/?directConnection=true&authMechanism=SCRAM-SHA-1"))
	if err != nil {
		log.Panic(err)
	}
	return client
}

type MongoDb interface {
	CreateUser(user User) error
}

func NewDB(dbName string, client *mongo.Client) *MongoDatabase {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Connect(ctx)
	if err != nil {
		log.Panic(err)
	}
	db := client.Database(dbName)
	return &MongoDatabase{
		Db: db,
	}
}

type MongoDatabase struct {
	Db *mongo.Database
}

func (db *MongoDatabase) FindByIds(ctx context.Context, collName string, ids []string, result interface{}) error {
	filter := bson.M{"_id": bson.M{"$in": IdsFromValidHex(ids)}}
	cur, err := db.Db.Collection(collName).Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cur.Close(ctx) // Close the cursor before returning from the function
	if err := cur.All(ctx, result); err != nil {
		return err
	}
	return nil
}

func IdsFromValidHex(ids []string) []primitive.ObjectID {
	objectIDs := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err == nil {
			objectIDs[i] = objectID
		}
	}
	return objectIDs
}
