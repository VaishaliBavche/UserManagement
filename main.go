package main

import (
	"UserManagement/apis"
	_ "UserManagement/apis/docs"
	"UserManagement/commons/apploggers"
	"UserManagement/configs"
	"UserManagement/internals/db"
	"UserManagement/internals/services"
	"context"

	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

// @User Management
// @version 1.0
// @description This is a sample API using Echo and Swagger.
// @contact.name API Support
// @contact.email support@example.com
// @host localhost:3000
// @BasePath /
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
	e.GET("/users/:id", userController.GetUserById)
	e.DELETE("/users/:id", userController.DeleteUserById)
	e.POST("/users", userController.CreateUser)

	// Swagger UI route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	logger.Infof("starting http server on localhost:%v", configs.AppConfig.HttpPort)
	e.Logger.Fatal(e.Start(":" + configs.AppConfig.HttpPort))
}
