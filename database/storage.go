package database

import (
	"context"

	"github.com/Milad75Rasouli/IEC61850CRUD/model"
)

type Storage interface {
	AddSignal(string, model.Signal, context.Context) error
	RemoveSignal(string, string, context.Context) error
	AllSignals(string, context.Context) ([]model.Signal, error)
}
