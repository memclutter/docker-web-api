package main

import (
	"flag"
	"net/http"
	"log"
)

var dist = flag.String("dist", "/var/www/html", "Angular frontend dist files")
var addr = flag.String("addr", ":8080", "Http address and port")
//var socket = flag.String("socket", "/var/run/docker.sock", "Docker unix-socket file")

func main() {
	flag.Parse()
	http.FileServer(http.Dir(*dist))
	http.Handle("/api/", http.HandlerFunc(api))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("Listen and serve:", err)
	}
}

func api(w http.ResponseWriter, req *http.Request) {
	log.Println("Handle request")
}
