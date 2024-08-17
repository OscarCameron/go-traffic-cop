package main

import (
	"log"
	"net/http"
)

func main() {

	servers := []Server{
		addServer("https://www.google.com/"),
		addServer("https://www.duckduckgo.com/"),
		addServer("https://www.youtube.com/"),
	}

	loadB := newLoadBalancer("8080", servers)

	handler := func(rw http.ResponseWriter, req *http.Request) {
		loadB.serveProxy(rw, req)
	}

	http.HandleFunc("/", handler)

	log.Printf("Welcome to the load balancer! Running on port %s", loadB.port)

	http.ListenAndServe(":"+loadB.port, nil)
}
