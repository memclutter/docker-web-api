package main

import (
	"net/http"
	"log"
)

func apiImages(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Printf("Access to images list")
}

func apiImagesDelete(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Printf("Access to image remove")
}
