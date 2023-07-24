package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var FileCollection *mongo.Collection

func ConnectToDB() error{
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client,err := mongo.Connect(context.Background(),clientOptions)
	if err != nil {
		return err
	}
	FileCollection = client.Database("test").Collection("file")
	return nil
	
}