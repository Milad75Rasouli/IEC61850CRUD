package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	db *mongo.Collection
}

func NewMongo(string)
