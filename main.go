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
package main

import (
	"api-gateway-knative-docker/config"
	"api-gateway-knative-docker/cors"
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
		http.HandleFunc(route.Path, proxy.HandleRequest(route, cors.GetCorsStore().Get()))
	}

	go docker.CheckContainersActive()
	go docker.CheckContainersToStop()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
