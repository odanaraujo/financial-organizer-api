package routers

import (
	userHandler "github.com/odanraujo/financial-organizer-api/internal/handler/users"
	"github.com/odanraujo/financial-organizer-api/internal/repository/users"
	usecase "github.com/odanraujo/financial-organizer-api/internal/usecase/users"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserInitDependecies(mongoDB *mongo.Database) userHandler.UserHandler {
	userRepo := users.NewUserRepository(mongoDB)
	userUsercase := usecase.NewUserUsecase(userRepo)
	userController := userHandler.NewUserHandler(userUsercase)

	return userController
}
