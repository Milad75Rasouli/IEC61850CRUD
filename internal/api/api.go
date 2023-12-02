package api

import (
	"net/http"

	"github.com/Milad75Rasouli/IEC61850CRUD/model"
	"github.com/Milad75Rasouli/IEC61850CRUD/utils"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	Endpoint string
}

func (a *ApiServer) Run() error {
	route := mux.NewRouter()
	route.HandleFunc("/IEC61850", utils.ErrorHandler(a.IEC61850))
	http.Handle("/", route)
	http.ListenAndServe(a.Endpoint, nil)

	return nil
}

func (a *ApiServer) IEC61850(w http.ResponseWriter, r *http.Request) error {

	err := utils.WriteJSON(w, http.StatusOK, model.IEC61850{Key: "voltage", Value: 310})
	if err != nil {
		return err
	}
	return nil
}

func NewApiServer(endpoint string) *ApiServer {
	return &ApiServer{
		Endpoint: endpoint,
	}
}
