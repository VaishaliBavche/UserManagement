package apis

import (
	"UserManagement/internals/db"
	"UserManagement/internals/services"
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
