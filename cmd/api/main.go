package main

import (
	"github.com/gin-gonic/gin"
	config "github.com/odanraujo/financial-organizer-api/config/env"
	"github.com/odanraujo/financial-organizer-api/infrastructure/database/mongodb"
	"github.com/odanraujo/financial-organizer-api/internal/handler/routers"
)

func main() {
	if _, err := config.LoadingConfig("."); err != nil {
		return
		//todo implement error and log here
	}

	database, err := mongodb.NewMongoDBConnection()

	if err != nil {
		return //todo implements log and error
	}
	userController := routers.UserInitDependecies(database)
	r := gin.Default()
	routers.InitUserRouters(&r.RouterGroup, userController)
}
