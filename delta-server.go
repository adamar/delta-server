package main

import (
	"github.com/adamar/delta-server/controllers"
	"github.com/adamar/delta-server/delta"
	"os"
)

func main() {

	go delta.RPCServer()

	deltaServer := controllers.Router()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3300"
	}

	deltaServer.Run(":" + port)

}
