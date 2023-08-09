package docker

import (
	"api-gateway-knative-docker/config"
	"context"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"log"
	"net/http"
	"sync"
	"time"
)

var containerMutex = &sync.Mutex{}

// StartContainer Funcionalidade de iniciar um container
func StartContainer(route config.Route) (bool, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return false, err
	}

	if !isContainerRunning(route.Service) {
		containerMutex.Lock()
		if !isContainerRunning(route.Service) {
			if err := cli.ContainerStart(ctx, route.Service, types.ContainerStartOptions{}); err != nil {
				defer containerMutex.Unlock()
				return false, err
			}

			log.Println("Iniciando container:", route.Service)

			//CheckContainersActive()

			// Verificar o healthcheck do container
			if !checkHealth(route) {
				return false, errors.New("healthcheck falhou para o container " + route.Service)
			}
			defer containerMutex.Unlock()
		}
	}

	return true, nil
}

// StopContainer Funcionalidade de parar um container
func StopContainer(containerID string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Println("Erro ao criar cliente Docker:", err)
		return
	}

	log.Println("Realizando stop no container:", containerID)
	containerMutex.Lock()

	err = cli.ContainerStop(ctx, containerID, container.StopOptions{}) // nil é um valor temporário, ajuste conforme necessário
	if err != nil {
		log.Println("Erro ao parar o container:", err)
	}

	//CheckContainersActive()

	defer containerMutex.Unlock()
}

// Verifica se o container está em execução
func isContainerRunning(service string) bool {
	ctx := context.Background()
	cli, _ := client.NewClientWithOpts(client.FromEnv)
	container, exists := GetContainerStore().Get(service)
	if !exists {
		return false
	}

	_, err := cli.ContainerInspect(ctx, container.ID)

	return err == nil
}

func checkHealth(route config.Route) bool {
	client := &http.Client{
		Timeout: time.Duration(route.RetryDelay) * time.Second, // Defina um timeout adequado
	}

	for i := 0; i < route.Retry; i++ {
		resp, err := client.Get("http://host.docker.internal:" + route.Port + route.HealthPath)
		if err == nil && resp.StatusCode == http.StatusOK {
			return true
		}

		// Se não for a última tentativa, aguarde antes de tentar novamente
		if i < route.Retry-1 {
			time.Sleep(time.Duration(route.RetryDelay) * time.Second)
		}
	}

	return false
}
