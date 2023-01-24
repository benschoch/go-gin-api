package db

import (
	"context"
	"log"
	"recipes-core-api/pkg/env"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB = ConnectDB()

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(env.LoadMongoENV()))
	if err != nil {
		log.Fatal(err)
	}

	dbConnectionTimeout := 10 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), dbConnectionTimeout)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to mongo: established")

	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("api").Collection(collectionName)

	return collection
}

func DropDB() {
	err := DB.Database("api").Drop(context.TODO())
	if err != nil {
		log.Println(err)
	}
}
