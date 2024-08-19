package services

import "UserManagement/internals/db"

type EventService interface {
	SaveUser() error
}

type eservice struct {
	dbservice db.DbService
}

func NewUserService(dbservice db.DbService) EventService {
	return &eservice{
		dbservice: dbservice,
	}
}

func (e *eservice) SaveUser() error {
	return nil
}
