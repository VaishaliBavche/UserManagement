package db

import "UserManagement/commons"

type udbservice struct {
	ucollection commons.DatabaseCollection
}

type DbService interface {
	SaveUser() error
}

func NewUserDbService(dbclient commons.DatabaseClient) DbService {
	return &udbservice{
		ucollection: dbclient.Collection("users"),
	}
}

func (u *udbservice) SaveUser() error {
	return nil
}
