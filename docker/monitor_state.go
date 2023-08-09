package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
	"strings"
	"sync"
	"time"
)

var updateContainerMutex sync.Mutex

func CheckContainersActive() {
	for {
		syncContainersState()
		time.Sleep(5 * time.Second)
	}
}

func syncContainersState() {
	updateContainerMutex.Lock()
	defer updateContainerMutex.Unlock()

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Println("Erro ao criar cliente Docker:", err)
		return
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		log.Println("Erro ao listar containers:", err)
		return
	}

	currentContainers := make(map[string]Container)
	for _, container := range containers {
		for _, name := range container.Names {
			// Criando um novo contêiner do tipo docker.Container
			newContainer := Container{
				ID:         container.ID,
				Service:    strings.ReplaceAll(name, "/", ""),
				LastAccess: time.Now(),
			}

			currentContainers[container.ID] = newContainer
		}
	}

	activeContainers := GetContainerStore().GetAll()

	// Atualizando o ContainerStore para refletir o estado real
	for _, storedContainer := range activeContainers {
		if _, exists := currentContainers[storedContainer.ID]; !exists {
			// Se o contêiner armazenado não estiver na lista de contêineres ativos, remova-o
			GetContainerStore().Remove(storedContainer.ID)
		}
	}

	for _, currentContainer := range currentContainers {
		if _, exists := activeContainers[currentContainer.ID]; !exists {
			// Se o contêiner ativo não estiver no containerStore, adicione-o
			GetContainerStore().Add(currentContainer)
		}
	}
}
