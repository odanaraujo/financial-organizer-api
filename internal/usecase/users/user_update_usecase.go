package users

import (
	"context"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	"github.com/odanraujo/financial-organizer-api/infrastructure/logger"
	entity "github.com/odanraujo/financial-organizer-api/internal/entity/users"
)

func (use *userUsecase) UpdateUserForCPFOrEmail(ctx context.Context, CPFOrEmail string, user entity.CreateUser) (entity.CreateUser, *excp.Exception) {

	logger.Info("init UpdateUserForCPFOrEmail in usecase")

	userForUpdate := entity.CreateUser{}
	userForUpdate.CPF = CPFOrEmail
	userForUpdate.Email = CPFOrEmail

	if userForUpdate.IsValidCPF() || userForUpdate.IsValidEmail() {
		if result, ex := use.repo.GetUserForCPF(ctx, userForUpdate.FormatCPF(CPFOrEmail)); ex == nil && result.CPF != "" {
			user, ex = use.repo.UpdateUserForCPF(ctx, CPFOrEmail, user)
			return user, nil
		}
		if result, ex := use.repo.GetUserForEmail(ctx, CPFOrEmail); ex == nil && result.Email != "" {
			user, ex = use.repo.UpdateUserForEmail(ctx, CPFOrEmail, user)
			return user, nil
		}
	}

	return entity.CreateUser{}, excp.NotFoundException("cpf or email not found")
}
