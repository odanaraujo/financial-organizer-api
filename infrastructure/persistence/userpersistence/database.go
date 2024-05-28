package userpersistence

import (
	"context"
	"database/sql"
	"github.com/odanraujo/financial-organizer-api/internal/domain/entity/userentity"
	"github.com/odanraujo/financial-organizer-api/internal/domain/repository/userrepository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) userrepository.UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) CreateUser(ctx context.Context, user userentity.User) error {
	//TODO implements save in database
	return nil
}
