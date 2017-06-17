package main

import (
	"flag"
	"net/http"
	"strings"
	"os"
	"log"

	"github.com/dimfeld/httptreemux"
)

var dist = flag.String("dist", "/var/www/html/", "Angular frontend dist folder")
var domain = flag.String("domain", "localhost", "Domain name")
var port = flag.String("port", "8080", "Expose port")

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

func main() {
	flag.Parse()

	log.Printf("Starting docker web ui server")
	log.Printf("Listening for HTTP connections at: http://%v:%v", *domain, *port)

	router := httptreemux.New()
	router.GET("/", indexHandler)
	router.GET("/*", staticHandler)

	if err := http.ListenAndServe((*domain) + ":" + (*port), router); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}