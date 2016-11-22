package controllers

import (
	"github.com/adamar/delta-server/models"
	"github.com/unrolled/render"
	"net/http"
	"log"
)

func StatsPage(w http.ResponseWriter, req *http.Request, template string, r *render.Render) {

	user := models.Event{}

	DBu.LogMode(true)
	DBu.Debug().Last(&user)

	log.Println(user)

	//log.Printf(user)

	//dbi.Limit(5).Find(&user)

	r.HTML(w, http.StatusOK, template, nil)
}
