package service

import (
	"context"
	"net/http"
	"time"

	"github.com/Milad75Rasouli/IEC61850CRUD/database"
	"github.com/Milad75Rasouli/IEC61850CRUD/model"
	"github.com/Milad75Rasouli/IEC61850CRUD/utils"
)

const (
	iec61850Collection = "IEC61850"
)

type IEC61850 struct {
	db      database.Storage
	ctxTime int
}

func NewIEC61850(d *database.Storage, ctxTime int) *IEC61850 {
	return &IEC61850{
		db:      *d,
		ctxTime: ctxTime,
	}
}

func (i *IEC61850) AllActions(w http.ResponseWriter, r *http.Request) error {

	var err error
	switch r.Method {
	case "POST":
		err = i.AddSignal(w, r)
	case "GET":
		err = i.AllSignals(w, r)
	case "DELETE":
		err = i.RemoveSignal(w, r)
	default:
		err = utils.WriteJSON(w, http.StatusBadRequest, model.ApiError{Error: r.Method})
	}

	if err != nil {
		return err
	}

	return nil
}

func (i *IEC61850) AddSignal(w http.ResponseWriter, r *http.Request) error {
	utils.WriteJSON(w, http.StatusOK, model.Signal{Key: "Post"})

	return nil
}

func (i *IEC61850) RemoveSignal(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (i *IEC61850) AllSignals(w http.ResponseWriter, r *http.Request) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(i.ctxTime)*time.Millisecond)
	defer cancel()
	data, err := i.db.AllSignals(iec61850Collection, ctx)
	if err != nil {
		return err
	}
	utils.WriteJSON(w, http.StatusOK, data)

	return nil
}
