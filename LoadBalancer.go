package main

import (
	"fmt"
	"net/http"
)

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func newLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func (lb *LoadBalancer) getNextAvailableServer() Server {
	serverIndex := lb.roundRobinCount % len(lb.servers)
	server := lb.servers[serverIndex]

	for !server.IsAlive() {
		lb.roundRobinCount++
		serverIndex = lb.roundRobinCount % len(lb.servers)
		server = lb.servers[serverIndex]
	}

	lb.roundRobinCount++

	return server
}

func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer()

	fmt.Printf("forwarding request to %s \n", targetServer.Address())

	targetServer.Serve(rw, req)
}
