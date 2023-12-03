package service

import (
	"github.com/Milad75Rasouli/IEC61850CRUD/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type IEC61850 struct {
	Collection *mongo.Collection
}

func (i *IEC61850) AddSignal(d model.Signal) error {
	return nil
}

func (i *IEC61850) RemoveSignal(d model.Signal) error {
	return nil
}

func (i *IEC61850) AllSignals() ([]model.Signal, error) {
	r := make([]model.Signal, 0)
	return r, nil
}
