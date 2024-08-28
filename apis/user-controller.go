package apis

import (
	"UserManagement/commons"
	"UserManagement/commons/apploggers"
	"UserManagement/internals/db"
	"UserManagement/internals/models"
	"UserManagement/internals/services"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ucontroller struct {
	dbservice db.DbService
	eservice  services.EventService
}

func NewUserController(dbservice db.DbService, eservice services.EventService) ucontroller {
	return ucontroller{
		dbservice: dbservice,
		eservice:  eservice,
	}
}

func (u *ucontroller) GetUsers(c echo.Context) error {
	lcontext, logger := apploggers.GetLoggerFromEcho(c)
	logger.Info("Executing Get All Users")
	users, eerror := u.eservice.GetUsers(lcontext)
	if eerror != nil {
		logger.Error(eerror)
		return eerror
	}
	logger.Infof("Executed GetUsers, users %s", commons.PrintStruct(users))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"total": len(users),
		"users": users,
	})
}

func (u *ucontroller) CreateUser(c echo.Context) error {
	lcontext, logger := apploggers.GetLoggerFromEcho(c)
	logger.Info("Executing CreateUser")
	name := "test_user" + uuid.New().String()
	serror := u.eservice.CreateUser(lcontext, &models.User{
		Name:     name,
		Type:     "Customer",
		Email:    name + "@gmail.com",
		IsActive: false,
	})
	if serror != nil {
		logger.Error(serror)
		return serror
	}
	logger.Info("Executed CreateUser")
	return c.NoContent(http.StatusCreated)
}
