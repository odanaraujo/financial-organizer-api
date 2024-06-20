package users

import (
	"github.com/gin-gonic/gin"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	"github.com/odanraujo/financial-organizer-api/infrastructure/logger"
	response "github.com/odanraujo/financial-organizer-api/internal/dto/response/users"
	"net/http"
)

func (u *userHandler) GetUserFotCPFOrEmail(ctx *gin.Context) {
	logger.Info("init request GetUserForCPF")

	cpfOrEmail := ctx.Param("cpfOrEmail")

	if cpfOrEmail == "" {
		ex := excp.NotFoundException("CPF or E-mail is required")
		logger.Error("CPF or E-mail is required", ex)
		ctx.JSON(ex.Code, ex)
		return
	}

	result, err := u.usecase.GetUserForCPF(ctx, cpfOrEmail)

	if err != nil {
		logger.Error("exception when searching for user", err)
		ctx.JSON(err.Code, err)
		return
	}

	resp := response.ConverterDomainToResponse(result)

	ctx.JSON(http.StatusOK, resp)
}
