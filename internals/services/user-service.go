package services

import (
	"UserManagement/commons/apploggers"
	"UserManagement/internals/db"
	dbmodel "UserManagement/internals/db/models"
	"UserManagement/internals/models"
	"context"
	"encoding/json"
)

type EventService interface {
	GetUserById(context context.Context, userId string) (*models.User, error)
	GetUsers(context context.Context) ([]*models.User, error)
	CreateUser(context context.Context, user *models.User) error
}

type eservice struct {
	dbservice db.DbService
}

func NewUserEventService(dbservice db.DbService) EventService {
	return &eservice{
		dbservice: dbservice,
	}
}

func (e *eservice) GetUserById(context context.Context, userId string) (*models.User, error) {
	logger := apploggers.GetLoggerWithCorrelationid(context)
	logger.Infof("Executing GetUserById, userId: %s", userId)
	user, dberror := e.dbservice.GetUserById(context, userId)
	if dberror != nil {
		logger.Error(dberror)
		return nil, dberror
	}
	logger.Infof("Executed GetUserById, userId: %s", userId)
	return user, nil
}

func (e *eservice) GetUsers(context context.Context) ([]*models.User, error) {
	logger := apploggers.GetLoggerWithCorrelationid(context)
	logger.Infof("Executing GetUsers...")
	users, dberror := e.dbservice.GetUsers(context)
	if dberror != nil {
		logger.Error(dberror)
		return nil, dberror
	}
	logger.Infof("Executed GetUsers, users: %d", len(users))
	return users, nil
}

func (e *eservice) CreateUser(context context.Context, user *models.User) error {
	logger := apploggers.GetLoggerWithCorrelationid(context)
	logger.Infof("Executing CreateUser...")
	var userSchema *dbmodel.UserSchema
	pbyes, _ := json.Marshal(user)
	uerror := json.Unmarshal(pbyes, &userSchema)
	if uerror != nil {
		logger.Error(uerror.Error())
		return uerror
	}
	userId, dberror := e.dbservice.SaveUser(context, userSchema)
	if dberror != nil {
		logger.Error(dberror)
		return dberror
	}
	logger.Infof("Executed CreateUser, userId: %v", userId)
	return nil
}
