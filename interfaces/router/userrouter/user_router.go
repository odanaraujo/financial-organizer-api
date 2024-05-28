package userrouter

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/odanraujo/financial-organizer-api/infrastructure/persistence/userpersistence"
	"github.com/odanraujo/financial-organizer-api/interfaces/controller/usercontroller"
	"github.com/odanraujo/financial-organizer-api/internal/usercase/userusecase"
)

func NewRouter(db *sql.DB) *gin.Engine {
	userRepo := userpersistence.NewUserRepository(db)
	userUsercase := userusecase.UserUsecase(userRepo)
	userController := usercontroller.NewUserController(userUsercase)

	r := gin.Default()
	r.POST("/users", userController.CreateUser)
	return r
}
