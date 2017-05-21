
package controllers

import (
	"net/http"
	"github.com/adamar/delta-server/models"
	"encoding/json"
	"github.com/gorilla/mux"
	)




func RecentEvents(w http.ResponseWriter, req *http.Request) {

        ev := []models.Event{}

        DBu.LogMode(true)
        DBu.Debug().Order("time_stamp desc").Limit(20).Find(&ev)

        data, _ := json.Marshal(ev)
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.Write(data)

}


func SingleEvent(w http.ResponseWriter, req *http.Request) {

        event := models.Event{}
	vars := mux.Vars(req)

	DBu.Where("uuid = ?", vars["eventId"]).Find(&event)

        data, _ := json.Marshal(event)
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.Write(data)

}



