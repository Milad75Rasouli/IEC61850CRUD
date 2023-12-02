package service

import "github.com/Milad75Rasouli/IEC61850CRUD/model"

type IEC61850 struct {
}

func (i *IEC61850) AddSignal(d model.IEC61850) error {
	return nil
}

func (i *IEC61850) RemoveSignal(d model.IEC61850) error {
	return nil
}

func (i *IEC61850) AllSignals() ([]model.IEC61850, error) {
	r := make([]model.IEC61850, 0)
	return r, nil
}
