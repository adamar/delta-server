
package controllers

import (
	"net/http"
	"github.com/adamar/delta-server/models"
	"encoding/json"
	)



func LatestEvent(w http.ResponseWriter, req *http.Request) {

        ev := models.Event{}

        DBu.LogMode(true)
        DBu.Debug().Last(&ev)

        data, _ := json.Marshal(ev)
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.Write(data)

}


func LatestEvents(w http.ResponseWriter, req *http.Request) {

        ev := []models.Event{}

        DBu.LogMode(true)
        DBu.Debug().Limit(2).Find(&ev)

        data, _ := json.Marshal(ev)
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.Write(data)

}





