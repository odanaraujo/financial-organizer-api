package users

import (
	"github.com/gin-gonic/gin"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	"github.com/odanraujo/financial-organizer-api/infrastructure/logger"
	dto "github.com/odanraujo/financial-organizer-api/internal/dto/request/users/update"
	response "github.com/odanraujo/financial-organizer-api/internal/dto/response/users"
	"go.uber.org/zap"
	"net/http"
)

func (u *userHandler) UpdateUser(ctx *gin.Context) {
	logger.Info("init request UpdateUser")

	request := dto.UserRequest{}

	cpfOrEmail := ctx.Param("cpfOrEmail")
	if cpfOrEmail == "" {
		ex := excp.NotFoundException("CPF or E-mail is required")
		logger.Error("CPF or E-mail is required", ex)
		ctx.JSON(ex.Code, ex)
		return
	}

	if ex := ctx.ShouldBindJSON(&request); ex != nil {
		logger.Error("error trying validate to user", ex, zap.String(
			"Journey", "UpdateUser"))
		errMessage := excp.InternalServerException(ex.Error())
		ctx.JSON(errMessage.Code, errMessage)
		return
	}

	userDatabase, ex := u.usecase.GetUserForCPFOrEmail(ctx, cpfOrEmail)

	if ex != nil {
		logger.Error("error trying validate to user", ex, zap.String(
			"Journey", "UpdateUser"))
		errMessage := excp.InternalServerException(ex.Error())
		ctx.JSON(errMessage.Code, errMessage)
		return
	}

	userDatabase = request.FillModel(userDatabase)

	result, ex := u.usecase.UpdateUserForCPFOrEmail(ctx, cpfOrEmail, userDatabase)

	if ex != nil {
		logger.Error("error when update for user in the database", ex, zap.String(
			"Journey", "UpdateUser"))

		ctx.JSON(ex.Code, ex)
		return
	}

	ctx.JSON(http.StatusOK, response.ConverterDomainToResponse(result))
}
