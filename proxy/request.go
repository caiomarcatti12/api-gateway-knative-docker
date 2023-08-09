package proxy

import (
	"api-gateway-knative-docker/config"
	"api-gateway-knative-docker/docker"
	"net/http"
	"net/url"
	"strings"
)

func HandleRequest(route config.Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, err := docker.StartContainer(route)
		if err != nil {
			http.Error(w, "Error starting container", http.StatusInternalServerError)
			return
		}

		serviceURL := &url.URL{
			Scheme: "http",
			Host:   "host.docker.internal:" + route.Port,
		}

		// Strip the route path from the request
		strippedPath := stripRoutePath(r.URL.Path, route.Path)
		r.URL.Path = strippedPath

		proxyToService(serviceURL)(w, r)
	}
}

func stripRoutePath(requestPath, routePath string) string {
	return strings.TrimPrefix(requestPath, routePath)
}
