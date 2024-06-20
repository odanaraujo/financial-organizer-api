package users

import (
	"context"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	entity "github.com/odanraujo/financial-organizer-api/internal/entity/users"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database: database}
}

type userRepository struct {
	database *mongo.Database
}

//go:generate mockgen -destination=mocks/UserRepository_mock.go -package=mocks github.com/odanraujo/financial-organizer-api/internal/repository/users UserRepository
type UserRepository interface {
	CreateUser(ctx context.Context, user entity.CreateUser) (entity.CreateUser, *excp.Exception)
	GetUserForCPF(ctx context.Context, cpf string) (entity.CreateUser, *excp.Exception)
	GetUserForEmail(ctx context.Context, email string) (entity.CreateUser, *excp.Exception)
	UpdateUserForCPF(ctx context.Context, cpf string, user entity.CreateUser) (entity.CreateUser, *excp.Exception)
	UpdateUserForEmail(ctx context.Context, email string, user entity.CreateUser) (entity.CreateUser, *excp.Exception)
}
