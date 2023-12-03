package api

import (
	"context"
	"net/http"
	"time"

	"github.com/Milad75Rasouli/IEC61850CRUD/database"
	"github.com/Milad75Rasouli/IEC61850CRUD/internal/service"
	"github.com/Milad75Rasouli/IEC61850CRUD/utils"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	Endpoint           string
	DB                 database.Storage
	iec61850Controller *service.IEC61850
}

func (a *ApiServer) Run() error {
	route := mux.NewRouter()

	a.iec61850Controller = service.NewIEC61850(&a.DB, 100)

	route.HandleFunc("/IEC61850", utils.ErrorHandler(a.iec61850Controller.AllActions))
	http.Handle("/", route)
	http.ListenAndServe(a.Endpoint, nil)

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
