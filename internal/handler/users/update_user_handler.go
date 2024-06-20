package users

import (
	"github.com/gin-gonic/gin"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	"github.com/odanraujo/financial-organizer-api/infrastructure/logger"
	dto "github.com/odanraujo/financial-organizer-api/internal/dto/request/users"
	response "github.com/odanraujo/financial-organizer-api/internal/dto/response/users"
	entity "github.com/odanraujo/financial-organizer-api/internal/entity/users"
	"go.uber.org/zap"
	"net/http"
)

func (u *userHandler) UpdateUser(ctx *gin.Context) {
	logger.Info("init request UpdateUser")

	request := dto.User{}

	cpfOrEmail := ctx.Param("cpfOrEmail")
	if cpfOrEmail == "" {
		ex := excp.NotFoundException("CPF or E-mail is required")
		logger.Error("CPF or E-mail is required", ex)
		ctx.JSON(ex.Code, ex)
		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error("error trying validate to user", err, zap.String(
			"Journey", "UpdateUser"))
		errMessage := excp.InternalServerException(err.Error())
		ctx.JSON(errMessage.Code, errMessage)
		return
	}

	userDomain := entity.NewUpdateUser(request)

	result, ex := u.usecase.UpdateUserForCPFOrEmail(ctx, cpfOrEmail, userDomain)

	if ex != nil {
		logger.Error("error when update for user in the database", ex, zap.String(
			"Journey", "UpdateUser"))

		ctx.JSON(ex.Code, ex)
		return
	}

	ctx.JSON(http.StatusOK, response.ConverterDomainToResponse(result))
}
