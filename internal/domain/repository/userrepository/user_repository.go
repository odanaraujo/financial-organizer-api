package userrepository

import (
	"context"
	"github.com/odanraujo/financial-organizer-api/internal/domain/entity/userentity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, entity userentity.User) error
}
