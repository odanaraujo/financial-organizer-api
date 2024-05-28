package usercontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/odanraujo/financial-organizer-api/internal/usercase/userusecase"
)

type userController struct {
	userCase userusecase.UserUsecase
}

func NewUserController(userCase userusecase.UserUsecase) *userController {
	return &userController{userCase: userCase}
}

func (usecase *userController) CreateUser(ctx *gin.Context) {
	//TODO implements user's controller
}
