package main

import (
	"github.com/Milad75Rasouli/IEC61850CRUD/database"
	"github.com/Milad75Rasouli/IEC61850CRUD/internal/api"
)

func main() {
	conn := database.MongoConnection{
		Port: 27017,
		IP:   "127.0.0.1",
		User: "root",
		Pass: "1234qwer",
	}

	server := api.NewApiServer(":5000")
	server.Run()
}
