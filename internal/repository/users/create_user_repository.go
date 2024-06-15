package users

import (
	"context"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	"github.com/odanraujo/financial-organizer-api/infrastructure/logger"
	entity "github.com/odanraujo/financial-organizer-api/internal/entity/users"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (repo *userRepository) CreateUser(ctx context.Context, user entity.CreateUser) (entity.CreateUser, *excp.Exception) {

	logger.Info("init CreateUser repository", zap.String("Journey", "CreateUser"))

	collectionName := viper.GetString(MONGODB_USER_COLLECTION)
	collection := repo.database.Collection(collectionName)

	result, ex := collection.InsertOne(ctx, user)

	if ex != nil {
		return entity.CreateUser{}, excp.BadRequestException("error when insert user")
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid.Hex()
		logger.Info("Inserted user ID")
	} else {
		logger.Info("Failed to convert InsertedID to ObjectID")
	}

	return user, nil
}
