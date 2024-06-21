package users

import (
	"context"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	"github.com/odanraujo/financial-organizer-api/infrastructure/logger"
	entity "github.com/odanraujo/financial-organizer-api/internal/entity/users"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (repo *userRepository) UpdateUserForCPF(ctx context.Context, cpf string, user entity.CreateUser) (entity.CreateUser, *excp.Exception) {
	logger.Info("init UpdateUserForCPF repository", zap.String("Journey", "UpdateUserForCPF"))

	collectionName := viper.GetString(MONGODB_USER_COLLECTION)
	collection := repo.database.Collection(collectionName)

	filter := bson.D{{Key: "cpf", Value: cpf}}
	update := bson.D{{Key: "$set", Value: user}}

	_, ex := collection.UpdateOne(ctx, filter, update)

	if ex != nil {
		return entity.CreateUser{}, excp.BadRequestException(ex.Error())
	}

	return user, nil
}

func (repo *userRepository) UpdateUserForEmail(ctx context.Context, email string, user entity.CreateUser) (entity.CreateUser, *excp.Exception) {
	logger.Info("init UpdateUserForEmail repository", zap.String("Journey", "UpdateUserForEmail"))

	collectionName := viper.GetString(MONGODB_USER_COLLECTION)
	collection := repo.database.Collection(collectionName)

	filter := bson.D{{Key: "email", Value: email}}
	update := bson.D{{Key: "$set", Value: user}}

	_, ex := collection.UpdateOne(ctx, filter, update)

	if ex != nil {
		return entity.CreateUser{}, excp.BadRequestException(ex.Error())
	}

	return user, nil
}
