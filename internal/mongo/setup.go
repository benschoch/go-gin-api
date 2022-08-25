package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"recipes-core-api/api/v1/models"
	"recipes-core-api/internal/env"
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

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to mongo: established")

	//initDemoData(*client)

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

func initDemoData(client mongo.Client) {
	log.Println("Dropping database...")

	err := client.Database("api").Drop(context.TODO())
	if err != nil {
		log.Println(err)
	}

	log.Println("Loading demo data...")

	ingredientCollection := client.Database("api").Collection("Ingredient")
	ingredientArray := []string{"salt", "pepper", "pasta", "meat", "oil"}
	for _, s := range ingredientArray {
		newIngredient := models.Ingredient{Id: primitive.NewObjectID(), Singular: s, Plural: s}
		_, err := ingredientCollection.InsertOne(context.TODO(), newIngredient)
		if err != nil {
			log.Println(err)
		}
	}

	unitCollection := client.Database("api").Collection("Unit")
	unitArray := []string{"g", "kg", "l", "ml", "handful"}
	for _, s := range unitArray {
		newUnit := models.Unit{Id: primitive.NewObjectID(), Singular: s, Plural: s}
		_, err := unitCollection.InsertOne(context.TODO(), newUnit)
		if err != nil {
			log.Println(err)
		}
	}
}
