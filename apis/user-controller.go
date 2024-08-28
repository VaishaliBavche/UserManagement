package apis

import (
	"UserManagement/commons/apploggers"
	"UserManagement/internals/db"
	"UserManagement/internals/services"
	"net/http"

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
	users, dberror := u.dbservice.GetUsers(lcontext)
	if dberror != nil {
		logger.Error(dberror)
		return dberror
	}
	logger.Info("Executed GetUsers, users %v", users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"total": len(users),
		"users": users,
	})
}

func (u *ucontroller) CreateUser(c echo.Context) error {
	lcontext, logger := apploggers.GetLoggerFromEcho(c)
	logger.Info("Executing CreateUser")
	dberror := u.dbservice.SaveUser(lcontext)
	if dberror != nil {
		logger.Error(dberror)
		return dberror
	}
	logger.Info("Executed CreateUser")
	return c.NoContent(http.StatusCreated)
}
