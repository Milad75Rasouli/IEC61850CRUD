package main

import (
	"log"

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

	server, err := api.NewApiServer(":5000", conn)
	if err != nil {
		log.Fatalln(err)
	}

	err = server.Run()
	if err != nil {
		log.Fatalln(err)
	}

}
