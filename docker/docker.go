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
	"api-gateway-knative-docker/config"
	"api-gateway-knative-docker/docker/container_store"
	"context"
	"errors"
	"log"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var (
	mutexes      = make(map[string]*sync.Mutex)
	mutexesGuard = &sync.Mutex{} // Guard para proteger o acesso ao mapa de mutexes
	once         sync.Once
)

// getDockerClient garante que apenas uma instância do cliente Docker seja criada (singleton).
func getDockerClient() (*client.Client, error) {
	var err error
	once.Do(func() {
		dockerClientInstance, err = client.NewClientWithOpts(client.FromEnv)
	})
	return dockerClientInstance, err
}

// getMutexForService retorna o mutex associado a um serviço, criando um novo se necessário.
func getMutexForService(service string) *sync.Mutex {
	mutexesGuard.Lock()
	defer mutexesGuard.Unlock()

	if _, exists := mutexes[service]; !exists {
		log.Printf("Criando novo mutex para o serviço: %s", service)
		mutexes[service] = &sync.Mutex{}
	}
	return mutexes[service]
}

// StartContainer Funcionalidade de iniciar um container
func StartContainer(route config.Route) (bool, error) {
	if route.ContainerName == "" {
		log.Println("Nenhum serviço associado à rota, ignorando start do container.")
		return true, nil
	}

	ctx := context.Background()
	cli, err := getDockerClient()

	if err != nil {
		log.Printf("Erro ao criar cliente Docker: %v", err)
		return false, err
	}

	log.Printf("Iniciando processo para o container do serviço: %s", route.ContainerName)

	containerService, exists := container_store.GetByContainerName(route.ContainerName)

	if !exists {
		log.Printf("Não foi possível encontrar o serviço para o container %s", route.ContainerName)
	}

	serviceMutex := getMutexForService(route.ContainerName)
	serviceMutex.Lock()
	defer serviceMutex.Unlock()

	log.Printf("Container para o serviço %s não está em execução. Tentando iniciar...", route.ContainerName)
	if err := cli.ContainerStart(ctx, containerService.ID, types.ContainerStartOptions{}); err != nil {
		log.Printf("Erro ao iniciar container para o serviço %s: %v", route.ContainerName, err)
		return false, err
	}

	log.Printf("Container iniciado para o serviço: %s", route.ContainerName)

	// Verificar o healthcheck do container
	if !checkHealth(route) {
		log.Printf("Healthcheck falhou para o container %s", route.ContainerName)
		return false, errors.New("healthcheck falhou para o container " + route.ContainerName)
	}

	log.Printf("Healthcheck bem-sucedido para o container: %s", route.ContainerName)


	log.Printf("Último acesso ao container do serviço %s atualizado.", route.ContainerName)
	return true, nil
}

// StopContainer Funcionalidade de parar um container
func StopContainer(containerID string) {
	ctx := context.Background()
	cli, err := getDockerClient()
	if err != nil {
		log.Printf("Erro ao criar cliente Docker: %v", err)
		return
	}

	log.Printf("Iniciando processo de stop para o container: %s", containerID)

	// Recupera o serviço associado ao containerID para obter o mutex correto
	service := getServiceForContainer(containerID)
	if service == "" {
		log.Printf("Erro ao encontrar o serviço associado ao container: %s", containerID)
		return
	}

	serviceMutex := getMutexForService(service)
	serviceMutex.Lock()
	defer serviceMutex.Unlock()

	log.Printf("Parando container: %s do serviço: %s", containerID, service)
	err = cli.ContainerStop(ctx, containerID, container.StopOptions{})
	if err != nil {
		log.Printf("Erro ao parar o container %s: %v", containerID, err)
	} else {
		log.Printf("Container %s parado com sucesso.", containerID)
	}
}
