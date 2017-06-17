package main

import (
	"net/http"
	"strings"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	http.ServeFile(w, r, *dist + "index.html")
}

// Handle Static Files (files containing a .extension)
// Send a not found error for any requests containing a .extension where the file does not exist
// Redirect any requests that were not for files to your to your SPA
func staticHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	requestPath := r.URL.Path
	fileSystemPath := *dist + r.URL.Path
	endURIPath := strings.Split(requestPath, "/")[len(strings.Split(requestPath, "/"))-1]
	splitPath := strings.Split(endURIPath, ".")
	if len(splitPath) > 1 {
		if f, err := os.Stat(fileSystemPath); err == nil && !f.IsDir() {
			http.ServeFile(w, r, fileSystemPath)
			return
		}
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, *dist+"index.html")
}
