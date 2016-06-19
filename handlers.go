package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func PutTrack(resp http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var track Track
	err := decoder.Decode(&track)
	if err != nil {
		fmt.Printf("Failed to put track. Error= %s", err.Error())
	}
	WriteTrack(track)
	write(nil, resp, req)
}

func GetTracks(resp http.ResponseWriter, req *http.Request) {
	tracks := ParseTracks()

	js, err := json.Marshal(tracks)
	if err != nil {
		fmt.Printf("Failed to marshal events %s\n", err.Error())
	}
	write(js, resp, req)
}

func GetTrack(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	event := ParseTrack(id)

	js, err := json.Marshal(event)
	if err != nil {
		fmt.Printf("Failed to marshal event %s\n", err.Error())
	}

	write(js, resp, req)
}

func write(toWrite []byte, resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "application/json")
	resp.WriteHeader(200)
	resp.Write(toWrite)
}
