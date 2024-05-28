package userusecase

import (
	"context"
	"github.com/odanraujo/financial-organizer-api/internal/domain/entity/userentity"
	"github.com/odanraujo/financial-organizer-api/internal/domain/repository/userrepository"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, entity userentity.User) error
}

type useUsecase struct {
	userRepo userrepository.UserRepository
}

func NewUserUsecase(userRepo userrepository.UserRepository) *useUsecase {
	return &useUsecase{userRepo: userRepo}
}

func (usecase *useUsecase) CreateUser(ctx context.Context, entity userentity.User) error {
	if err := usecase.userRepo.CreateUser(ctx, entity); err != nil {
		return err
	}
	return nil
}
