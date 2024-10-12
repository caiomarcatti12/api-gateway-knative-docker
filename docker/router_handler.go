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
	"log"
	"net/http"
	"time"
)

func checkHealth(route config.Route) bool {
	client := &http.Client{
		Timeout: time.Duration(route.RetryDelay) * time.Second, // Defina um timeout adequado
	}

	log.Printf("Verificando healthcheck para o serviço: %s", route.ContainerName)
	for i := 0; i < route.Retry; i++ {
		resp, err := client.Get("http://host.docker.internal:" + route.Port + route.HealthPath)
		if err == nil && resp.StatusCode == http.StatusOK {
			log.Printf("Healthcheck bem-sucedido para o serviço: %s na tentativa %d", route.ContainerName, i+1)
			return true
		}

		log.Printf("Tentativa %d de healthcheck falhou para o serviço: %s, erro: %v", i+1, route.ContainerName, err)

		// Se não for a última tentativa, aguarde antes de tentar novamente
		if i < route.Retry-1 {
			time.Sleep(time.Duration(route.RetryDelay) * time.Second)
		}
	}

	log.Printf("Healthcheck falhou para o serviço: %s após %d tentativas.", route.ContainerName, route.Retry)
	return false
}

// getServiceForContainer é um placeholder para obter o serviço associado ao containerID
// (Você precisaria implementar isso com base no seu store)
func getServiceForContainer(containerID string) string {
	container, exists := container_store.GetByID(containerID)
	if exists {
		return container.ContainerName
	}
	log.Printf("Não foi possível encontrar o serviço para o container %s", containerID)
	return ""
}
