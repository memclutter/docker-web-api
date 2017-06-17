package main

import (
	"net/http"
	"log"
)

func apiImagesList(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Printf("Access to images list")
}

func apiImagesRemove(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Printf("Access to image remove")
}
