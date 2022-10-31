package mapper

import (
	"go-morning/model/dbmodel"
	"go-morning/model/request"
)

func Request_to_dbmodel(r request.UpsertDani) dbmodel.Dani {
	return dbmodel.Dani{
		Name: r.Name,
	}
}
