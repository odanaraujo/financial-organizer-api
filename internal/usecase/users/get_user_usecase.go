package users

import (
	"context"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	"github.com/odanraujo/financial-organizer-api/infrastructure/logger"
	entity "github.com/odanraujo/financial-organizer-api/internal/entity/users"
)

func (use *userUsecase) GetUser(ctx context.Context, ID string) (entity.CreateUser, *excp.Exception) {
	logger.Info("init GetUser in usecase")

	result, ex := use.repo.GetUser(ctx, ID)

	if ex != nil {
		return entity.CreateUser{}, ex
	}

	result.ID = ID

	return result, nil
}
