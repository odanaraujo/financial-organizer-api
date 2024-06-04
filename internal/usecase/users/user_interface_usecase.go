package users

import (
	"context"
	userRequest "github.com/odanraujo/financial-organizer-api/internal/dto/users"
	"github.com/odanraujo/financial-organizer-api/internal/repository/users"
	response "github.com/odanraujo/financial-organizer-api/internal/response/users"
)

type userUsecase struct {
	repo users.UserRepository
}

func NewUserUsecase(repo users.UserRepository) UserUsercase {
	return &userUsecase{repo: repo}
}

//go:generate mockgen -destination=mocks/UserUsercase_mock.go -package=mocks github.com/odanraujo/financial-organizer-api/internal/usecase/users UserUsercase
type UserUsercase interface {
	CreateUser(ctx context.Context, user userRequest.User) (response.User, error)
}
