package main

import (
	"flag"
	"net/http"
	"log"

	"github.com/dimfeld/httptreemux"
)

var dist = flag.String("dist", "/var/www/html/", "Angular frontend dist folder")
var domain = flag.String("domain", "localhost", "Domain name")
var port = flag.String("port", "8080", "Expose port")

func main() {
	flag.Parse()

	log.Printf("Starting docker web ui server")
	log.Printf("Listening for HTTP connections at: http://%v:%v", *domain, *port)

	router := httptreemux.New()
	router.GET("/", indexHandler)
	router.GET("/*", staticHandler)

	router.GET("/api/images", apiImagesList)
	router.DELETE("/api/images/:id", apiImagesRemove)

	router.GET("/api/containers", apiContainersList)
	router.DELETE("/api/containers/:id", apiContainersRemove)

	if err := http.ListenAndServe((*domain) + ":" + (*port), router); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}