package delta

import (
	"encoding/gob"
	"github.com/adamar/delta-server/models"
	"github.com/valyala/gorpc"
	"log"
)

func RPCServer() {

	gob.RegisterName("Event", models.Event{})
	gob.RegisterName("Response", models.Response{})

	s := gorpc.Server{

		// Accept clients on this TCP address.
		Addr: "0.0.0.0:12345",

		Handler: DoIt,
	}

	if err := s.Serve(); err != nil {
		log.Fatalf("Cannot start rpc server: %s", err)
	}

}

func DoIt(clientAddr string, request interface{}) interface{} {

	//log.Printf("Obtained request %+v from the client %s\n", request, clientAddr)

	blerg := request.(models.Event)

	insertEvent(&blerg)

	return models.Response{Result: "OK"}

}

func insertEvent(evt *models.Event) error {

	DBi.LogMode(true)
	DBi.Debug().Create(&evt)
	//DBi.Create(&evt)

	return nil

}
