package users

import (
	"context"
	userRequest "github.com/odanraujo/financial-organizer-api/internal/dto/users"
	userResponse "github.com/odanraujo/financial-organizer-api/internal/response/users"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database: database}
}

type userRepository struct {
	database *mongo.Database
}

//go:generate mockgen -destination=mocks/UserRepository_mock.go -package=mocks github.com/odanraujo/financial-organizer-api/internal/repository/users UserRepository
type UserRepository interface {
	CreateUser(ctx context.Context, user userRequest.User) (userResponse.User, error)
}
