package main

import (
	"api-gateway-knative-docker/config"
	"api-gateway-knative-docker/docker"
	"api-gateway-knative-docker/proxy"
	"log"
	"net/http"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		return
	}

	for _, route := range config.GetRouteStore().GetAll() {
		http.HandleFunc(route.Path, proxy.HandleRequest(route))
	}

	go docker.CheckContainersActive()
	go docker.CheckContainersToStop()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
