package users

import (
	"context"
	userRequest "github.com/odanraujo/financial-organizer-api/internal/dto/users"
	userResponse "github.com/odanraujo/financial-organizer-api/internal/response/users"
)

func (repo *userRepository) CreateUser(ctx context.Context, address userRequest.Address) (userResponse.Address, error) {
	return userResponse.Address{}, nil
}
