package users

import (
	"context"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	"github.com/odanraujo/financial-organizer-api/infrastructure/logger"
	entity "github.com/odanraujo/financial-organizer-api/internal/entity/users"
)

func (use *userUsecase) GetUserFotCPFOrEmail(ctx context.Context, CPFOrEmail string) (entity.CreateUser, *excp.Exception) {
	logger.Info("init GetUserFotCPFOrEmail in usecase")

	result, ex := use.repo.GetUserFotCPFOrEmail(ctx, CPFOrEmail)

	if ex != nil {
		return entity.CreateUser{}, ex
	}

	return result, nil
}
