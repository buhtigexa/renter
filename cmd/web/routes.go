package main

import "net/http"

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/resource", app.showResource)
	mux.HandleFunc("/resource/create", app.createResource)
	return mux
}
