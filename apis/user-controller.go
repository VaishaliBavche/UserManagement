package apis

import (
	"UserManagement/commons"
	"UserManagement/commons/apploggers"
	"UserManagement/internals/db"
	"UserManagement/internals/models"
	"UserManagement/internals/services"
	"net/http"
	"strings"

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

func (u *ucontroller) GetUserById(c echo.Context) error {
	lcontext, logger := apploggers.GetLoggerFromEcho(c)
	userId := c.Param("id")
	logger.Infof("Executing GetUserById, userId: %s", userId)
	user, serror := u.eservice.GetUserById(lcontext, userId)
	if serror != nil {
		logger.Error(serror)
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse(serror.Error(), nil))
	}
	logger.Infof("Executed GetUserById, userId:%s, user %s", userId, commons.PrintStruct(user))
	return c.JSON(http.StatusOK, user)
}

func (u *ucontroller) DeleteUserById(c echo.Context) error {
	lcontext, logger := apploggers.GetLoggerFromEcho(c)
	userId := c.Param("id")
	logger.Infof("Executing DeleteUserById, userId: %s", userId)
	serror := u.eservice.DeleteUserById(lcontext, userId)
	if serror != nil {
		logger.Error(serror)
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse(serror.Error(), nil))
	}
	logger.Infof("Executed DeleteUserById, userId: %s", userId)
	return c.NoContent(http.StatusNoContent)
}

func (u *ucontroller) GetUsers(c echo.Context) error {
	lcontext, logger := apploggers.GetLoggerFromEcho(c)
	logger.Info("Executing Get All Users")
	users, serror := u.eservice.GetUsers(lcontext)
	if serror != nil {
		logger.Error(serror)
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse(serror.Error(), nil))
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
	var user *models.User
	err := c.Bind(&user)
	if err != nil || user == nil {
		logger.Error("invalid request payload")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("invalid request payload", nil))
	}

	if len(strings.TrimSpace(user.Name)) == 0 {
		logger.Error("'name' is required")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("'name' is required", nil))
	}

	if len(strings.TrimSpace(user.Email)) == 0 {
		logger.Error("'email' is required")
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("'email' is required", nil))
	}
	Id, serror := u.eservice.CreateUser(lcontext, user)
	if serror != nil {
		logger.Error(serror)
		return c.JSON(http.StatusBadRequest, commons.ApiErrorResponse(serror.Error(), nil))
	}
	logger.Info("Executed CreateUser")
	return c.JSON(http.StatusCreated, map[string]string{
		"id": Id,
	})
}
