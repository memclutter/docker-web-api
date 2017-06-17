package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/dimfeld/httptreemux"
)

var dist = flag.String("dist", "/var/www/html/", "Angular frontend dist folder")
var domain = flag.String("domain", "localhost", "Domain name")
var port = flag.String("port", "8080", "Expose port")

func main() {
	flag.Parse()

	log.Printf("Starting docker web ui server")
	log.Printf("Serve angular frontend from %v", dist)
	log.Printf("Listening for HTTP connections at: http://%v:%v", *domain, *port)

	router := httptreemux.New()
	router.GET("/", indexHandler)
	router.GET("/*", staticHandler)

	router.GET("/api/images", apiImages)
	router.DELETE("/api/images/:id", apiImagesDelete)

	router.GET("/api/containers", apiContainers)
	router.POST("/api/containers/:id/start", apiContainersStart)
	router.POST("/api/containers/:id/stop", apiContainersStop)
	router.POST("/api/containers/:id/restart", apiContainersRestart)
	router.DELETE("/api/containers/:id", apiContainersDelete)

	router.GET("/api/networks", apiNetworks)
	router.DELETE("/api/networks/:id", apiNetworksDelete)

	router.GET("/api/volumes", apiVolumes)
	router.DELETE("/api/volumes/:id", apiVolumesDelete)

	if err := http.ListenAndServe((*domain)+":"+(*port), router); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}
