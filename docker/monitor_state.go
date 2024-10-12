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
	"api-gateway-knative-docker/docker/container_store"
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	updateContainerMutex sync.Mutex
	dockerClientInstance *client.Client
)

// CheckContainersActive inicia o processo contínuo de verificação dos containers.
func CheckContainersActive() {
	for {
		syncContainersState()
		time.Sleep(5 * time.Second)
	}
}

// syncContainersState é o processo principal que sincroniza o estado dos containers.
func syncContainersState() {
	updateContainerMutex.Lock()
	defer updateContainerMutex.Unlock()

	cli, err := getDockerClient()
	if err != nil {
		log.Println("Erro ao obter cliente Docker:", err)
		return
	}

	containers, err := listAllContainers(cli)
	if err != nil {
		log.Println("Erro ao listar containers:", err)
		return
	}

	currentContainers := mapContainers(containers)
	activeContainers := container_store.GetAll()

	removeMissingContainers(activeContainers, currentContainers)
	updateOrAddContainers(activeContainers, currentContainers)
}

// listAllContainers lista todos os containers, incluindo os parados.
func listAllContainers(cli *client.Client) ([]types.Container, error) {
	ctx := context.Background()
	return cli.ContainerList(ctx, types.ContainerListOptions{All: true})
}

// mapContainers cria um mapa dos containers atuais com suas informações relevantes.
func mapContainers(containers []types.Container) map[string]container_store.Container {
	currentContainers := make(map[string]container_store.Container)

	for _, container := range containers {
		for _, name := range container.Names {
			newContainer := createContainerObject(container, name)
			currentContainers[container.ID] = newContainer
		}
	}
	return currentContainers
}

// createContainerObject cria uma instância de Container com base nos dados fornecidos.
func createContainerObject(container types.Container, name string) container_store.Container {
	return container_store.Container{
		ID:            container.ID,
		ContainerName: strings.ReplaceAll(name, "/", ""),
		LastAccess:    time.Now(),
		IsActive:      container.State == "running",
	}
}

// removeMissingContainers remove containers que não estão mais presentes no host.
func removeMissingContainers(activeContainers, currentContainers map[string]container_store.Container) {
	for containerID, storedContainer := range activeContainers {
		if _, exists := currentContainers[containerID]; !exists {
			container_store.Remove(containerID)
			log.Printf("Removido contêiner: %s (%s)", storedContainer.ContainerName, storedContainer.ID)
		}
	}
}

// updateOrAddContainers adiciona ou atualiza containers no store.
func updateOrAddContainers(activeContainers, currentContainers map[string]container_store.Container) {
	for _, currentContainer := range currentContainers {
		if storedContainer, exists := activeContainers[currentContainer.ID]; exists {
			updateContainerIfChanged(storedContainer, currentContainer)
		} else {
			addNewContainer(currentContainer)
		}
	}
}

// updateContainerIfChanged atualiza um container se houver mudança no status.
func updateContainerIfChanged(storedContainer, currentContainer container_store.Container) {
	if storedContainer.IsActive != currentContainer.IsActive {
		storedContainer.IsActive = currentContainer.IsActive

		container_store.Update(storedContainer)

		log.Printf("Atualizado contêiner: %s (%s) - IsActive: %v",
			storedContainer.ContainerName, storedContainer.ID, storedContainer.IsActive)
	}
}

// addNewContainer adiciona um novo container ao store.
func addNewContainer(currentContainer container_store.Container) {
	container_store.Add(currentContainer)
	log.Printf("Adicionado novo contêiner: %s (%s)", currentContainer.ContainerName, currentContainer.ID)
}
