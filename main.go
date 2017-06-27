package main

import (
	"io"
	"log"
	"net/http"
)

const version string = "3.1.4"

// VersionHandler handles incoming requests to /version
// and just returns a simple version number
func versionHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, version)
}

func meetupHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "HELLO AWS MEETUP")
}

func main() {
	log.Printf("Listening on port 8000...")
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/aws", meetupHandler)
	http.ListenAndServe(":8000", nil)
}
