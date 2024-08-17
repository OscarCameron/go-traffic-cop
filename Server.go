package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server interface {
	// Address returns the address to access the server.
	Address() string

	// Returns true if the server is alive.
	IsAlive() bool

	// Use this server to process the request.
	Serve(rw http.ResponseWriter, r *http.Request)
}

type server struct {
	address string
	proxy   *httputil.ReverseProxy
}

func (s *server) Address() string {
	return s.address
}

func (s *server) IsAlive() bool {
	return true
}

func (s *server) Serve(rw http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(rw, r)
}

func addServer(address string) *server {
	serverURL, err := url.Parse(address)
	if err != nil {
		log.Fatalf("%s occured while parsing %s", err, serverURL)
	}

	return &server{
		address: address,
		proxy:   httputil.NewSingleHostReverseProxy(serverURL),
	}

}
