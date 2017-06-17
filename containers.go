package main

import (
	"net/http"
	"log"
)

func apiContainersList(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Printf("Access to containers list")
}

func apiContainersRemove(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Printf("Access to containers remove")
}
