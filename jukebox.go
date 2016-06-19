package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/trackById/{id:[0-9]+}", GetTrack)
	r.HandleFunc("/tracks/", GetTracks)
	r.HandleFunc("/putTrack/", PutTrack).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9080", r))

}
