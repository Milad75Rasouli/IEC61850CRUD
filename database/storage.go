package database

import "github.com/Milad75Rasouli/IEC61850CRUD/model"

type Storage interface {
	CreateSignal(model.IEC61850) error
	RemoveSignal(model.IEC61850) error
	AllSignals(model.IEC61850) error
}
