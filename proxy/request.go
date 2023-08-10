/*
 * Copyright 2023 Caio Matheus Marcatti Calimério
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
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
