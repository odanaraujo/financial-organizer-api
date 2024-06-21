package users

import (
	"context"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	"github.com/odanraujo/financial-organizer-api/infrastructure/logger"
	entity "github.com/odanraujo/financial-organizer-api/internal/entity/users"
)

func (use *userUsecase) GetUserForCPFOrEmail(ctx context.Context, CPFOrEmail string) (entity.CreateUser, *excp.Exception) {
	logger.Info("init GetUserForCPFOrEmail in usecase")

	user := entity.CreateUser{}
	user.CPF = CPFOrEmail
	user.Email = CPFOrEmail

	var ex *excp.Exception

	if user.IsValidCPF() {

		user, ex = use.repo.GetUserForCPF(ctx, user.FormatCPF(CPFOrEmail))
		if ex == nil {
			return user, nil
		}
	}

	if user.IsValidEmail() {
		user, ex = use.repo.GetUserForEmail(ctx, CPFOrEmail)
		if ex == nil {
			return user, nil
		}
	}

	return entity.CreateUser{}, excp.BadRequestException("CPF or email parameter is mandatory")
}
