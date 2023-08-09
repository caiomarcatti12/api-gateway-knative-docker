package docker

import (
	"sync"
	"time"
)

var containerStore *ContainerStore
var once sync.Once

// Expondo o ContainerStore para outras partes do programa
func GetContainerStore() *ContainerStore {
	once.Do(func() {
		containerStore = newContainerStore()
	})
	return containerStore
}

// tornando privado para garantir que somente GetContainerStore() possa criar uma instância
func newContainerStore() *ContainerStore {
	return &ContainerStore{
		containers: make(map[string]Container),
	}
}

func (cs *ContainerStore) Add(container Container) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.containers[container.ID] = container
}

func (cs *ContainerStore) Remove(containerID string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	delete(cs.containers, containerID)
}

func (cs *ContainerStore) Get(containerID string) (Container, bool) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	container, exists := cs.containers[containerID]
	return container, exists
}

func (cs *ContainerStore) GetByService(serviceName string) (*Container, bool) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	for _, container := range cs.containers {
		if container.Service == serviceName {
			return &container, true
		}
	}
	return nil, false
}

func (cs *ContainerStore) UpdateAccessTime(containerID string, accessTime time.Time) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	if container, exists := cs.containers[containerID]; exists {
		container.LastAccess = accessTime
		cs.containers[containerID] = container
	}
}

func (cs *ContainerStore) GetAll() map[string]Container {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	// Retornando uma cópia para evitar mutações externas
	copiedContainers := make(map[string]Container, len(cs.containers))
	for id, container := range cs.containers {
		copiedContainers[id] = container
	}
	return copiedContainers
}
