package users

import (
	"github.com/gin-gonic/gin"
	"github.com/odanraujo/financial-organizer-api/internal/usecase/users"
)

func NewUserHandler(usecase users.UserUsercase) UserHandler {
	return &userHandler{usecase: usecase}
}

type userHandler struct {
	usecase users.UserUsercase
}

type UserHandler interface {
	CreateUser(ctx *gin.Context)
	GetUserFotCPFOrEmail(ctx *gin.Context)
}
