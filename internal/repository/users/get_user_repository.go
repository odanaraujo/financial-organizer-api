package users

import (
	"context"
	"fmt"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	"github.com/odanraujo/financial-organizer-api/infrastructure/logger"
	entity "github.com/odanraujo/financial-organizer-api/internal/entity/users"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (repo *userRepository) GetUserFotCPFOrEmail(ctx context.Context, CPFOrEmail string) (entity.CreateUser, *excp.Exception) {

	logger.Info("init GetUserFotCPFOrEmail repository", zap.String("Journey", "GetUserFotCPFOrEmail"))

	collectionName := viper.GetString(MONGODB_USER_COLLECTION)
	collection := repo.database.Collection(collectionName)

	userentity := entity.CreateUser{}

	filter := bson.D{{Key: "cpf", Value: CPFOrEmail}}

	ex := collection.FindOne(ctx, filter).Decode(&userentity)
	if ex == mongo.ErrNoDocuments {
		filter = bson.D{{Key: "email", Value: CPFOrEmail}}
		ex = collection.FindOne(ctx, filter).Decode(&userentity)
		if ex == mongo.ErrNoDocuments {
			errMessage := fmt.Sprintf("user not found with this params %s", CPFOrEmail)
			return userentity, excp.NotFoundException(errMessage)
		}
	}

	return userentity, nil
}
