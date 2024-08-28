package db

import (
	"UserManagement/commons"
	"UserManagement/commons/appdb"
	"UserManagement/commons/apploggers"
	"UserManagement/configs"
	dbmodel "UserManagement/internals/db/models"
	"UserManagement/internals/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type udbservice struct {
	ucollection appdb.DatabaseCollection
}

type DbService interface {
	GetUsers(ctx context.Context) ([]*models.User, error)
	SaveUser(ctx context.Context, user *dbmodel.UserSchema) (string, error)
}

func NewUserDbService(dbclient appdb.DatabaseClient) DbService {
	return &udbservice{
		ucollection: dbclient.Collection(configs.MONGO_USERS_COLLECTION),
	}
}

func (u *udbservice) GetUsers(ctx context.Context) ([]*models.User, error) {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing GetUsers")

	// create users payload to find data from db
	var users []*models.User
	var filter = map[string]interface{}{}
	dbError := u.ucollection.Find(ctx, filter, &options.FindOptions{}, &users)
	if dbError != nil {
		logger.Error(dbError)
		return nil, dbError
	}
	logger.Infof("Executed GetUsers, users: %d", len(users))
	return users, nil
}

func (u *udbservice) SaveUser(ctx context.Context, user *dbmodel.UserSchema) (string, error) {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing SaveUser...")

	// insert user in db
	result, dbError := u.ucollection.InsertOne(ctx, user)
	if dbError != nil {
		logger.Error(dbError)
		return "", dbError
	}

	// Extract the inserted ID from the result
	id := result.InsertedID.(primitive.ObjectID).Hex()
	logger.Infof("Executed SaveUser, userid: %s", commons.PrintStruct(user))
	return id, nil
}
