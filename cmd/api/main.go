package main

import (
	"github.com/gin-gonic/gin"
	config "github.com/odanraujo/financial-organizer-api/config/env"
	"github.com/odanraujo/financial-organizer-api/infrastructure/database/mongodb"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	"github.com/odanraujo/financial-organizer-api/infrastructure/logger"
	"github.com/odanraujo/financial-organizer-api/internal/handler/routers"
)

func main() {
	if _, err := config.LoadingConfig("."); err != nil {
		return
		//todo implement error and log here
	}

	database, err := mongodb.NewMongoDBConnection()

	if err != nil {
		logger.Error("error connection in database", err)
		excp.InternalServerException("error connection in database")
		return
	}
	userController := routers.UserInitDependecies(database)
	r := gin.Default()
	routers.InitUserRouters(&r.RouterGroup, userController)

	if err := r.Run(); err != nil {
		logger.Error("error init server", err)
		return
	}
}
