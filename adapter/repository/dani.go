package repository

import (
	"context"
	"fmt"
	"go-morning/infrastructure/database/mongodb"
	"go-morning/model/dbmodel"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func GetMorningMessages() ([]dbmodel.Morning, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var mornings []dbmodel.Morning
	cursor, err := mongodb.GetCollection("mornings").Find(ctx, bson.M{})
	if err != nil {
		return []dbmodel.Morning{}, err
	}
	if err = cursor.All(ctx, &mornings); err != nil {
		return []dbmodel.Morning{}, err
	}
	defer cancel()

	fmt.Println(mornings)
	return mornings, nil
}
