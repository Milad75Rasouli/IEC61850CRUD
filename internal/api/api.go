package api

import (
	"context"
	"net/http"
	"time"

	"github.com/Milad75Rasouli/IEC61850CRUD/database"
	"github.com/Milad75Rasouli/IEC61850CRUD/model"
	"github.com/Milad75Rasouli/IEC61850CRUD/utils"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	Endpoint string
	DB       database.Storage
}

func (a *ApiServer) Run() error {
	route := mux.NewRouter()
	route.HandleFunc("/IEC61850", utils.ErrorHandler(a.IEC61850))
	http.Handle("/", route)
	http.ListenAndServe(a.Endpoint, nil)

	return nil
}

func (a *ApiServer) IEC61850(w http.ResponseWriter, r *http.Request) error {

	err := utils.WriteJSON(w, http.StatusOK, model.Signal{Key: "voltage", Value: 310})
	if err != nil {
		return err
	}
	return nil
}

func NewApiServer(endpoint string, dbconfig database.MongoConnection) (*ApiServer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	db, err := database.NewMongo(dbconfig, ctx)
	if err != nil {
		return nil, err
	}

	return &ApiServer{
		Endpoint: endpoint,
		DB:       db,
	}, nil
}
