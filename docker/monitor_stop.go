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
