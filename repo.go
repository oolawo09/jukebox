package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Track struct {
	Url  string `json:"url"`
	Id   string `json:"id"`
	Type string `json:"type"`
}

func ParseTracks() []Track {
	tracksFile, err := os.Open("./tracks.json")
	if err != nil {
		fmt.Printf("Couldn't open tracks.json %s\n", err.Error())
	}

	jsonParser := json.NewDecoder(tracksFile)
	var tracks []Track
	if err = jsonParser.Decode(&tracks); err != nil {
		fmt.Printf("parsing events file: %v \n", err.Error())
	}

	return tracks
}

func ParseTrack(id string) Track {
	tracks := ParseTracks()
	var track Track

	for _, t := range tracks {
		if t.Id == id {
			track = t
		}
	}
	return track
}

func WriteTrack(track Track) {
	tracks := ParseTracks()
	tracks = append(tracks, track)
	tracksToWrite, err_1 := json.Marshal(tracks)
	if err_1 != nil {
		fmt.Printf("Failed to write to file %s\n", err_1.Error())
	}
	err := ioutil.WriteFile("./tracks.json", tracksToWrite, 0644)
	if err != nil {
		fmt.Printf("Failed to write to file %s\n", err.Error())
	}
}
