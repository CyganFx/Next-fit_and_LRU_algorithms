package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	getRouter := r.Methods("GET").Subrouter()
	postRouter := r.Methods("POST").Subrouter()

	getRouter.HandleFunc("/", app.home)
	getRouter.HandleFunc("/LRU", app.LRUHome)
	postRouter.HandleFunc("/doNextFit", app.doNextFit)
	postRouter.HandleFunc("/LRU/doLRU", app.doLRU)
	return r
}
