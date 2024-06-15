package users

import (
	"github.com/gin-gonic/gin"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	"github.com/odanraujo/financial-organizer-api/infrastructure/logger"
	response "github.com/odanraujo/financial-organizer-api/internal/dto/response/users"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
)

func (u *userHandler) GetUser(ctx *gin.Context) {
	logger.Info("init request GetUser")

	id := ctx.Param("id")

	if id == "" {
		ex := excp.NotFoundException("id is required")
		logger.Error("id is required", ex)
		ctx.JSON(ex.Code, ex)
		return
	}

	if _, ex := primitive.ObjectIDFromHex(id); ex != nil {
		logger.Error("Error trying to validate id", ex, zap.String(
			"Journey", "GetUser"))
		errMessage := excp.BadRequestException("id is not valid")
		ctx.JSON(errMessage.Code, errMessage)
		return
	}

	result, err := u.usecase.GetUser(ctx, id)

	if err != nil {
		logger.Error("exception when searching for user", err)
		ctx.JSON(err.Code, err)
	}

	resp := response.ConverterDomainToResponse(result)

	ctx.JSON(http.StatusOK, resp)
}
