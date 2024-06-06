package users

import (
	"context"
	userRequest "github.com/odanraujo/financial-organizer-api/internal/dto/request/users"
	userResponse "github.com/odanraujo/financial-organizer-api/internal/dto/response/users"
)

func (repo *userRepository) CreateUser(ctx context.Context, user userRequest.User) (userResponse.User, error) {
	return userResponse.User{}, nil
}
