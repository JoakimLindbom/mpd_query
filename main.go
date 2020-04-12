package main

import (
	"fmt"
	"github.com/fhs/gompd/mpd"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const host = "192.168.1.18"

//type server struct{}

func getArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"artists": ["Art of Noise", "ARC", "TOOL", "A Perfect Circle"]}`))
}

func getArtist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"artist": "Art of Noise", "no of records": 7, "no of songs": 35}`))
}

func get(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var msg string = dial()

	w.Write([]byte(`{"message":"` + msg + `"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post was called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func dial() string {
	// Connect to MPD server
	conn, err := mpd.Dial("tcp", host+":6600")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	line := ""
	line1 := ""
	rtn := ""
	// Loop printing the current status of MPD.
	status, err := conn.Status()
	if err != nil {
		log.Fatalln(err)
	}
	song, err := conn.CurrentSong()
	if err != nil {
		log.Fatalln(err)
	}
	if status["state"] == "play" {
		line1 = fmt.Sprintf("%s - %s", song["Artist"], song["Title"])
	} else {
		line1 = fmt.Sprintf("State: %s", status["state"])
	}
	if line != line1 {
		line = line1
		fmt.Println(line)
	}
	rtn += line

	return rtn
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/artists/", getArtists).Methods(http.MethodGet)
	api.HandleFunc("/artists/{artist}", getArtist).Methods(http.MethodGet)
	api.HandleFunc("/", get).Methods(http.MethodGet)
	api.HandleFunc("/", post).Methods(http.MethodPost)
	api.HandleFunc("/", notFound)
	log.Fatal(http.ListenAndServe(":5555", r))
}
