package users

import (
	"context"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	entity "github.com/odanraujo/financial-organizer-api/internal/entity/users"
	"github.com/odanraujo/financial-organizer-api/internal/repository/users"
)

type userUsecase struct {
	repo users.UserRepository
}

func NewUserUsecase(repo users.UserRepository) UserUsercase {
	return &userUsecase{repo: repo}
}

//go:generate mockgen -destination=mocks/UserUsercase_mock.go -package=mocks github.com/odanraujo/financial-organizer-api/internal/usecase/users UserUsercase
type UserUsercase interface {
	CreateUser(ctx context.Context, user entity.CreateUser) (entity.CreateUser, *excp.Exception)
	GetUserForCPF(ctx context.Context, CPFOrEmail string) (entity.CreateUser, *excp.Exception)
}
