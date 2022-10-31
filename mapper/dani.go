package mapper

import (
	"go-morning/entity/dbmodel"
	"go-morning/entity/request"
)

func Request_to_dbmodel(r request.UpsertDani) dbmodel.Dani {
	return dbmodel.Dani{
		Name: r.Name,
	}
}
