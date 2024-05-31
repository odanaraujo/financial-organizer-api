package mongodb

import (
	"context"
	config "github.com/odanraujo/financial-organizer-api/config/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBConnection() (*mongo.Database, error) {
	mongoURI := config.Env.DatabaseURL
	mongoUser := config.Env.DatabaseUser

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return client.Database(mongoUser), nil
}
