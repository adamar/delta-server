package controllers

import (
	"github.com/adamar/delta-server/delta"
	"github.com/codegangsta/negroni"
	"github.com/unrolled/render"
	"net/http"
	"github.com/gorilla/mux"
)

var (
	rootServer *negroni.Negroni
	port       string
	DBu        = delta.SetupGormDB()
)

func Router() *negroni.Negroni {

	t := render.New(render.Options{
		Directory: "views",
	})

	n := negroni.Classic()
	mux := mux.NewRouter()



        ////////////////////////// API - v1 ////////////////////////////////



        //////////  events  //////////

	mux.HandleFunc("/api/v1/events/recent", RecentEvents)

	mux.HandleFunc("/api/v1/events/{eventId}", SingleEvent)


        //////////  hosts  //////////

	//mux.HandleFunc("/api/v1/hosts/recent", RecentHosts)

	//mux.HandleFunc("/api/v1/hosts/{hostId}", SingleHost)







        ////////////////////////// Pages ////////////////////////////////

   
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		SimplePage(w, r, "mainpage", t)
	})

	mux.HandleFunc("/recent-events", func(w http.ResponseWriter, r *http.Request) {
		SimplePage(w, r, "latest", t)
	})

	mux.HandleFunc("/event/{eventId}", func(w http.ResponseWriter, r *http.Request) {
		SimplePage(w, r, "event", t)
	})







	n.UseHandler(mux)

	return n

}
