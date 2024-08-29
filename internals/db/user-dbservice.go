package db

import (
	"UserManagement/commons"
	"UserManagement/commons/appdb"
	"UserManagement/commons/apploggers"
	"UserManagement/configs"
	dbmodel "UserManagement/internals/db/models"
	"UserManagement/internals/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type udbservice struct {
	ucollection appdb.DatabaseCollection
}

type DbService interface {
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
	SaveUser(ctx context.Context, user *dbmodel.UserSchema) (string, error)
}

func NewUserDbService(dbclient appdb.DatabaseClient) DbService {
	return &udbservice{
		ucollection: dbclient.Collection(configs.MONGO_USERS_COLLECTION),
	}
}

func (u *udbservice) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	logger := apploggers.GetLoggerWithCorrelationid(ctx)
	logger.Infof("Executing GetUserById, Id: %s", userId)
	// get object if=d from userid string
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, fmt.Errorf("invalid userId provided, userId: %s", userId)
	}
	var user *models.User
	var filter = bson.M{"_id": id}
	dbError := u.ucollection.FindOne(ctx, filter, &user)
	if dbError != nil {
		logger.Error(dbError)
		return nil, dbError
	}
	logger.Infof("Executed GetUserById, user: %s", commons.PrintStruct(user))
	return user, nil
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
