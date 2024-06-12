package users

import (
	"context"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	"github.com/odanraujo/financial-organizer-api/infrastructure/logger"
	entity "github.com/odanraujo/financial-organizer-api/internal/entity/users"
	"go.uber.org/zap"
)

func (use *userUsecase) CreateUser(ctx context.Context, userEntity entity.CreateUser) (entity.CreateUser, *excp.Exception) {
	logger.Info("init usecase CreateUser")

	if !userEntity.IsValidCPF() {
		ex := excp.BadRequestException("is not a valid CPF")
		logger.Error("is not a valid CPF", ex, zap.String("Journey", "CreateUser"))
		return entity.CreateUser{}, ex
	}

	userResult, ex := use.repo.CreateUser(ctx, userEntity)

	if ex != nil {
		logger.Error("unable to create database", ex)
		return entity.CreateUser{}, ex
	}
	
	return userResult, nil
}
