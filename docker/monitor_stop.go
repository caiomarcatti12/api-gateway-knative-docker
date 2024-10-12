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
package docker

import (
	"api-gateway-knative-docker/config" // Importar o pacote correto
	"api-gateway-knative-docker/docker/container_store"
	"sync"
	"time"
)

var (
	containerMonitorMutex sync.Mutex
)

// CheckContainersToStop inicia o processo contínuo de monitoramento e parada de containers inativos.
func CheckContainersToStop() {
	for {
		monitorAndStopContainers()
		time.Sleep(5 * time.Second)
	}
}

// monitorAndStopContainers monitora e para containers que estão inativos além do tempo limite.
func monitorAndStopContainers() {
	containerMonitorMutex.Lock()
	defer containerMonitorMutex.Unlock()

	now := time.Now()
	routeStore := config.GetRouteStore()
	containers := container_store.GetAll()

	for _, container := range containers {
		checkAndStopContainer(container, routeStore, now)
	}
}

// checkAndStopContainer verifica se o container deve ser parado com base no TTL.
func checkAndStopContainer(container container_store.Container, routeStore *config.RouteStore, now time.Time) {
	route, exists := routeStore.Get(container.ContainerName)
	if !exists {
		return
	}

	if isContainerExpired(container, route, now) {
		stopAndRemoveContainer(container)
	}
}

// isContainerExpired verifica se o container excedeu o tempo de inatividade permitido.
func isContainerExpired(container container_store.Container, route config.Route, now time.Time) bool {
	return now.Sub(container.LastAccess) > time.Duration(route.TTL)*time.Second && container.IsActive
}

// stopAndRemoveContainer para e remove o container da store.
func stopAndRemoveContainer(container container_store.Container) {
	StopContainer(container.ID)

	container.IsActive = false
	container_store.Update(container)
}
