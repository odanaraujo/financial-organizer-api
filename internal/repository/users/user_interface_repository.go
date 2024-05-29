package users

import (
	"context"
	"database/sql"
	userRequest "github.com/odanraujo/financial-organizer-api/internal/dto/users"
	userResponse "github.com/odanraujo/financial-organizer-api/internal/response/users"
)

func NewUserRepository(database *sql.DB) UserRepository {
	return &userRepository{database: database}
}

type userRepository struct {
	database *sql.DB
}

type UserRepository interface {
	CreateUser(ctx context.Context, address userRequest.Address) (userResponse.Address, error)
}
