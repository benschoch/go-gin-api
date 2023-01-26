package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	recipeCollection      = "recipes"
	courseCollection      = "courses"
	dietCollection        = "diets"
	ingredientCollection  = "ingredients"
	unitCollection        = "units"
	regionCollection      = "regions"
	foodTypeCollection    = "foodtypes"
	servingTypeCollection = "servingtypes"
)

type Connection struct {
	client    *mongo.Client
	defaultDB string
}

func NewConnection(config *Config) (*Connection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.DBURL))
	if err != nil {
		return nil, err
	}

	dbConnectionTimeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), dbConnectionTimeout)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &Connection{client: client, defaultDB: config.DBName}, nil
}

func (c *Connection) GetCollection(name string) *mongo.Collection {
	return c.selectDefaultDB().Collection(name)
}

func (c *Connection) GetRecipes() *mongo.Collection {
	return c.selectDefaultDB().Collection(recipeCollection)
}

func (c *Connection) GetCourses() *mongo.Collection {
	return c.selectDefaultDB().Collection(courseCollection)
}

func (c *Connection) GetDiets() *mongo.Collection {
	return c.selectDefaultDB().Collection(dietCollection)
}

func (c *Connection) GetIngredients() *mongo.Collection {
	return c.selectDefaultDB().Collection(ingredientCollection)
}

func (c *Connection) GetUnits() *mongo.Collection {
	return c.selectDefaultDB().Collection(unitCollection)
}

func (c *Connection) GetRegions() *mongo.Collection {
	return c.selectDefaultDB().Collection(regionCollection)
}

func (c *Connection) GetFoodTypes() *mongo.Collection {
	return c.selectDefaultDB().Collection(foodTypeCollection)
}

func (c *Connection) GetServingTypes() *mongo.Collection {
	return c.selectDefaultDB().Collection(servingTypeCollection)
}

func (c *Connection) selectDefaultDB() *mongo.Database {
	return c.client.Database(c.defaultDB)
}

func (c *Connection) DropDB() error {
	return c.selectDefaultDB().Drop(context.TODO())
}
