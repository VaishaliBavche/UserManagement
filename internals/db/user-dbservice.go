package db

import (
	"UserManagement/commons/appdb"
)

type udbservice struct {
	ucollection appdb.DatabaseCollection
}

type DbService interface {
	SaveUser() error
}

func NewUserDbService(dbclient appdb.DatabaseClient) DbService {
	return &udbservice{
		ucollection: dbclient.Collection("users"),
	}
}

func (u *udbservice) SaveUser() error {
	return nil
}
