package main

import (
	"github.com/adamar/delta-server/controllers"
	"github.com/adamar/delta-server/delta"
)

func main() {

	// Start Delta Server 
	go delta.RPCServer()

	// Start the Web Frontend
	controllers.RunFrontend()

}
