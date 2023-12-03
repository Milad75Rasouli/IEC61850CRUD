package database

import (
	"context"
	"fmt"

	"github.com/Milad75Rasouli/IEC61850CRUD/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	IP   string
	Port int
	User string
	Pass string
}
type Mongo struct {
	db *mongo.Database
}

func NewMongo(conn MongoConnection, ctx context.Context) (*Mongo, error) {
	strConn := fmt.Sprintf("mongodb://%s:%s@%s:%d", conn.User, conn.Pass, conn.IP, conn.Port)
	Option := options.Client().ApplyURI(strConn)
	dbConfig, err := mongo.Connect(ctx, Option)
	db := dbConfig.Database("protocol") //.Collection("protocols")

	if err != nil {
		return nil, err
	}
	return &Mongo{db: db}, nil
}

func (m *Mongo) AddSignal(collection string, s model.Signal, ctx context.Context) error {
	_, err := m.db.Collection(collection).InsertOne(ctx, s)
	return err
}

func (m *Mongo) RemoveSignal(collection string, k string, ctx context.Context) error {
	bsonKey := bson.D{{"key", k}}
	_, err := m.db.Collection(collection).DeleteOne(ctx, bsonKey)
	return err
}

func (m *Mongo) AllSignals(collection string, ctx context.Context) ([]model.Signal, error) {
	s := []model.Signal{}

	result, err := m.db.Collection(collection).Find(ctx, bson.D{})

	result.All(ctx, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
