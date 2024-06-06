package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/odanraujo/financial-organizer-api/internal/handler/users"
)

func InitUserRouters(r *gin.RouterGroup, handler users.UserHandler) {
	r.POST("/financial/createuser", handler.CreateUser)
}
