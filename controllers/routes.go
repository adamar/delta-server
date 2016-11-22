package controllers

import (
	"github.com/adamar/delta-server/delta"
	"github.com/codegangsta/negroni"
	"github.com/unrolled/render"
	"net/http"
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
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		SimplePage(w, r, "mainpage", t)
	})

	mux.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		StatsPage(w, r, "latest", t)
	})

	mux.HandleFunc("/api/last", func(w http.ResponseWriter, r *http.Request) {
		LatestEvent(w, r)
	})

	mux.HandleFunc("/api/latest", func(w http.ResponseWriter, r *http.Request) {
		LatestEvents(w, r)
	})


	n.UseHandler(mux)

	return n

}
