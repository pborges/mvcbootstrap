package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/pborges/mvc"
	"github.com/pborges/log"
	"github.com/pborges/mvcbootstrap/home"
)

const ListenAddress string = ":8080"

func main() {
	log.SetLogLevel(log.LevelDebug)
	mvc.CacheTemplates = false
	log.Info("starting")
	r := mux.NewRouter()

	homeController := home.RegisterController()
	r.HandleFunc("/", homeController.Index).Methods("GET")

	// static
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Warn("listening on", ListenAddress)
	log.Panic(http.ListenAndServe(ListenAddress, Log(r)))
}

func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithField("ip", r.RemoteAddr).WithField("url", r.URL).WithField("method", r.Method).Info("http request")
		h.ServeHTTP(w, r)
	})
}