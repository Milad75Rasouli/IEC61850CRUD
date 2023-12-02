package main

import "github.com/Milad75Rasouli/IEC61850CRUD/internal/api"

func main() {
	server := api.NewApiServer(":5000")

	server.Run()
}
