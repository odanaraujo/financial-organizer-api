package users

import (
	"github.com/gin-gonic/gin"
	"github.com/odanraujo/financial-organizer-api/infrastructure/logger"
	"github.com/odanraujo/financial-organizer-api/infrastructure/validator"
	dto "github.com/odanraujo/financial-organizer-api/internal/dto/request/users"
	response "github.com/odanraujo/financial-organizer-api/internal/dto/response/users"
	"github.com/odanraujo/financial-organizer-api/internal/entity/users"
	"go.uber.org/zap"
	"net/http"
)

func (u *userHandler) CreateUser(ctx *gin.Context) {
	logger.Info("init request createUser")
	var userRequest dto.User

	if ex := ctx.ShouldBindJSON(&userRequest); ex != nil {
		logger.Error("error trying to validate user info", ex)
		excp := validator.ValidateUserRequest(ex)
		ctx.JSON(excp.Code, excp)
		return
	}

	createUserDomain := users.NewCreateUser(userRequest)
	domainResult, ex := u.usecase.CreateUser(ctx, createUserDomain)

	if ex != nil {
		logger.Error("error the trying create user", ex, zap.String("Journey", "CreateUser"))
		ctx.JSON(ex.Code, ex)
		return
	}

	logger.Info("user created successfuly", zap.String("Journey", "CreateUser"))
	result := response.ConverterDomainToResponse(domainResult)
	ctx.JSON(http.StatusOK, result)
}
