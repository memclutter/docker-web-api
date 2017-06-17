package main

import (
	"net/http"
	"log"
)

func apiVolumes(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Printf("Access to volumes list")
}

func apiVolumesDelete(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Printf("Access to volumes delete")
}
