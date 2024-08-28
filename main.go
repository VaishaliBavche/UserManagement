package main

import (
	"UserManagement/apis"
	"UserManagement/commons/apploggers"
	"UserManagement/configs"
	"UserManagement/internals/db"
	"UserManagement/internals/services"
	"context"

	"github.com/labstack/echo/v4"
)

func main() {
	context, logger := apploggers.NewLoggerWithCorrelationid(context.Background(), "")
	err := configs.NewApplicationConfig(context)
	if err != nil {
		logger.Errorf("Error in Appconfig:", err)
	}

	dbservice := db.NewUserDbService(configs.AppConfig.DbClient)
	eventService := services.NewUserEventService(dbservice)

	// Echo instance
	e := echo.New()

	// user api Routes
	userController := apis.NewUserController(dbservice, eventService)
	e.GET("/users", userController.GetUsers)
	e.POST("/users", userController.CreateUser)

	// Start server
	logger.Infof("starting http server on localhost:%v", configs.AppConfig.HttpPort)
	e.Logger.Fatal(e.Start(":" + configs.AppConfig.HttpPort))
}
