package db

import (
	"UserManagement/commons/appdb"
	"UserManagement/commons/apploggers"
	"UserManagement/configs"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type udbservice struct {
	ucollection appdb.DatabaseCollection
}

type DbService interface {
	GetUsers(ctx context.Context) ([]map[string]interface{}, error)
	SaveUser(ctx context.Context) error
}

func NewUserDbService(dbclient appdb.DatabaseClient) DbService {
	return &udbservice{
		ucollection: dbclient.Collection(configs.MONGO_USERS_COLLECTION),
	}
}

func (u *udbservice) GetUsers(ctx context.Context) ([]map[string]interface{}, error) {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing GetUsers")
	var users = []map[string]interface{}{}
	var filter = map[string]interface{}{}
	dbError := u.ucollection.Find(ctx, filter, &options.FindOptions{}, &users)
	if dbError != nil {
		logger.Error(dbError)
		return nil, dbError
	}
	logger.Infof("Executed GetUsers, users: %d", len(users))
	return users, nil
}

func (u *udbservice) SaveUser(ctx context.Context) error {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	name := "test_user" + uuid.New().String()
	var user = map[string]interface{}{
		"name":  name,
		"user":  "Customer",
		"email": name + "@gmail.com",
	}
	logger.Infof("Executing SaveUser")
	_, dbError := u.ucollection.InsertOne(ctx, user)
	if dbError != nil {
		return dbError
	}
	logger.Infof("Executed SaveUser, user: %v", user)
	return nil
}
