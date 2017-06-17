package main

import (
	"net/http"
	"log"
)

func apiContainers(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Printf("Access to containers list")
}

func apiContainersDelete(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Printf("Access to containers remove")
}
