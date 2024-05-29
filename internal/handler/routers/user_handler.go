package routers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	userHandler "github.com/odanraujo/financial-organizer-api/internal/handler/users"
	"github.com/odanraujo/financial-organizer-api/internal/repository/users"
	usecase "github.com/odanraujo/financial-organizer-api/internal/usecase/users"
)

func NewRouter(db *sql.DB) *gin.Engine {
	userRepo := users.NewUserRepository(db)
	userUsercase := usecase.NewUserUsecase(userRepo)
	userController := userHandler.NewUserHandler(userUsercase)

	r := gin.Default()
	r.POST("/routers", userController.CreateUser)
	return r
}
