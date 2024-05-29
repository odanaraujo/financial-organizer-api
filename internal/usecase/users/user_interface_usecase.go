package users

import (
	"context"
	userRequest "github.com/odanraujo/financial-organizer-api/internal/dto/users"
	"github.com/odanraujo/financial-organizer-api/internal/repository/users"
	users2 "github.com/odanraujo/financial-organizer-api/internal/response/users"
)

type userUsecase struct {
	repo users.UserRepository
}

func NewUserUsecase(repo users.UserRepository) UserUsercase {
	return &userUsecase{repo: repo}
}

type UserUsercase interface {
	CreateUser(ctx context.Context, address userRequest.Address) (users2.Address, error)
}
