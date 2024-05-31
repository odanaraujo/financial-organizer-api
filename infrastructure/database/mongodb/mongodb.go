package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

const (
	MONGODB_URL     = "DATA_BASE_MONGO_URL"
	MONGODB_USER_DB = "MONGO_USERNAME"
)

func NewMongoDBConnection() (*mongo.Database, error) {
	mongoURI := os.Getenv(MONGODB_URL)
	mongoUser := os.Getenv(MONGODB_USER_DB)

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
