package repository

import (
	"context"
	"fmt"
	"go-morning/infrastructure/database/mongodb"
	"go-morning/model/dbmodel"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetEveningMessages() ([]dbmodel.Evening, error) {
	return getMessages[dbmodel.Evening]("evenings")
}

func GetMorningMessages() ([]dbmodel.Morning, error) {
	return getMessages[dbmodel.Morning]("mornings")
}

func getMessages[T any](collectionName string) ([]T, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var messages []T
	cursor, err := mongodb.GetCollection(collectionName).Find(ctx, bson.M{})
	if err != nil {
		return []T{}, err
	}
	if err = cursor.All(ctx, &messages); err != nil {
		return []T{}, err
	}
	defer cancel()

	fmt.Println(messages)
	return messages, nil
}
