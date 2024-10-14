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
	"fmt"
	"log"
	"net/http"
	"time"
)

// checkHealth realiza o healthcheck para uma rota específica usando Retry e Liveness Probe.
func checkHealth(route config.RouteConfig) bool {
	client := &http.Client{
		Timeout: time.Duration(route.Retry.Period) * time.Second,
	}

	// Extrair configuração do Liveness Probe
	liveness := route.LivenessProbe
	url := fmt.Sprintf("%s://%s:%d/%s", route.Backend.Protocol, route.Backend.Host, route.Backend.Port, route.LivenessProbe.Path)

	log.Printf("Verificando healthcheck para o serviço: %s", route.Backend.ContainerName)

	// Espera inicial definida no Liveness Probe
	if liveness.InitialDelaySeconds > 0 {
		log.Printf("Aguardando %d segundos antes do healthcheck inicial...", liveness.InitialDelaySeconds)
		time.Sleep(time.Duration(liveness.InitialDelaySeconds) * time.Second)
	}

	// Tentativas definidas no RetryConfig
	for attempt := 1; attempt <= route.Retry.Attempts; attempt++ {
		resp, err := client.Get(url)

		// Verificação de sucesso
		if err == nil && resp.StatusCode == http.StatusOK {
			log.Printf("Healthcheck bem-sucedido para %s na tentativa %d",
				route.Backend.ContainerName, attempt)
			return true
		}

		log.Printf("Tentativa %d falhou para %s, erro: %v",
			attempt, route.Backend.ContainerName, err)

		// Se não for a última tentativa, aguarde o período de retry
		if attempt < route.Retry.Attempts {
			log.Printf("Aguardando %d segundos antes da próxima tentativa...", route.Retry.Period)
			time.Sleep(time.Duration(route.Retry.Period) * time.Second)
		}
	}

	// Se todas as tentativas falharem, aguardar o período de tolerância para término (grace period)
	log.Printf("Healthcheck falhou para %s após %d tentativas. Aguardando %d segundos antes de finalizar...",
		route.Backend.ContainerName, route.Retry.Attempts, route.TTL)

	time.Sleep(time.Duration(route.TTL) * time.Second)
	return false
}
