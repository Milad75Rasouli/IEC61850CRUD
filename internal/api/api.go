package api

import (
	"fmt"
	"net/http"
)

type ApiServer struct {
	Endpoint string
}

func (a *ApiServer) Run() error {
	fmt.Println("Sub? Dude!")
	return nil
}

func (a *ApiServer) IEC61850(w *http.ResponseWriter, r *http.Request) error {

	return nil
}
