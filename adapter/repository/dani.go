package repository

import (
	"context"
	"go-morning/infrastructure/database/mongodb"
	"go-morning/model/dbmodel"
	"log"
	"time"
)

func SaveDani(dani dbmodel.Dani) {
	var morningsCollection = mongodb.GetCollection("mornings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := morningsCollection.InsertOne(ctx, dani)
	if err != nil {
		log.Panic("Error to save dani=", err)
	}
}
