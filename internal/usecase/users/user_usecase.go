package users

import (
	"context"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	entity "github.com/odanraujo/financial-organizer-api/internal/entity/users"
)

func (use *userUsecase) CreateUser(ctx context.Context, user entity.CreateUser) (entity.CreateUser, *excp.Exception) {
	return entity.CreateUser{}, nil
}
