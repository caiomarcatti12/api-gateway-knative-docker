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
	"sync"
	"time"
)

var containerMonitorMutex sync.Mutex

func CheckContainersToStop() {
	for {
		containerMonitorMutex.Lock()

		now := time.Now()
		routeStore := config.GetRouteStore()
		containers := GetContainerStore().GetAll()

		for _, container := range containers {
			route, exists := routeStore.Get(container.Service)
			if !exists {
				continue
			}

			if now.Sub(container.LastAccess) > time.Duration(route.TTL)*time.Second {
				StopContainer(container.ID)
				GetContainerStore().Remove(container.ID)
			}
		}

		containerMonitorMutex.Unlock()
		time.Sleep(5 * time.Second)
	}
}
