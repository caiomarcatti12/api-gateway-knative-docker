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
package config

// Configuração de um host específico
type HostConfig struct {
	Host   string        `yaml:"host"`   // Host para qual rotas serão configuradas
	CORS   CORSConfig    `yaml:"cors"`   // Configuração CORS específica para este host
	Routes []RouteConfig `yaml:"routes"` // Lista de rotas do host
}

// Configuração de uma rota específica
type RouteConfig struct {
	Path          string              `yaml:"path"`          // Caminho da rota
	StripPath     bool                `yaml:"stripPath"`     // Indica se o path deve ser removido
	TTL           int                 `yaml:"ttl"`           // Período de tolerância para término
	Backend       Backend             `yaml:"backend"`       // Configuração do backend
	Retry         RetryConfig         `yaml:"retry"`         // Configuração de tentativas (retry)
	LivenessProbe LivenessProbeConfig `yaml:"livenessProbe"` // Configuração do healthcheck
}

// Backend da rota
type Backend struct {
	Protocol      string `yaml:"protocol"`      // Protocolo (http ou https)
	Host          string `yaml:"host"`          // Host do backend
	Port          int    `yaml:"port"`          // Porta do backend
	ContainerName string `yaml:"containerName"` // Nome do container correspondente
}

// Configuração de retry para a rota
type RetryConfig struct {
	Attempts int `yaml:"attempts"` // Número de tentativas
	Period   int `yaml:"period"`   // Intervalo entre tentativas em segundos
}

// Configuração de liveness probe (healthcheck)
type LivenessProbeConfig struct {
	Path                string `yaml:"path"`                // Caminho do healthcheck
	SuccessThreshold    int    `yaml:"successThreshold"`    // Threshold de sucesso
	InitialDelaySeconds int    `yaml:"initialDelaySeconds"` // Tempo inicial antes do healthcheck
}
