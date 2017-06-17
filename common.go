package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HttpErrorResponse struct {
	Message string `json:"message"`
}

// Send http error as JSON object like as {"message": "Error occurred"} or raw string
func SendHttpError(code int, message string, w http.ResponseWriter) {
	w.WriteHeader(code)

	if body, err := json.Marshal(HttpErrorResponse{Message: message}); err != nil {
		w.Write([]byte(message))
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Write(body)
	}

	log.Printf("HTTP %v: %v", code, message)
}

// Get query string param, if not exists return defaultValue
func QueryGetOrDefaultValue(r *http.Request, key string, defaultValue string) string {
	if value := r.URL.Query().Get(key); value == "" {
		return defaultValue
	} else {
		return value
	}
}

// Send JSON data or send error
func SendJSONOrError(w http.ResponseWriter, data interface{}) {
	if body, err := json.Marshal(data); err != nil {
		SendHttpError(500, fmt.Sprintf("%v", err), w)
	} else {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}

// Send empty body
func SendOk(w http.ResponseWriter) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("[]"))
}
