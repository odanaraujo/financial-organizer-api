package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/odanraujo/financial-organizer-api/internal/handler/users"
)

func InitUserRouters(r *gin.RouterGroup, handler users.UserHandler) {
	r.POST("/financial", handler.CreateUser)
	r.GET("/financial/:id", handler.GetUser)
}
