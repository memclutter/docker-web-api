package main

import (
	"net/http"
	"log"
)

func apiNetworks(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Printf("Access to networks list")
}

func apiNetworksDelete(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Printf("Access to networks delete")
}
