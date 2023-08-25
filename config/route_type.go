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

import "api-gateway-knative-docker/cors"

type Route struct {
	Protocol   string           `yaml:"protocol"` // Protocolo (http ou https)
	Host       string           `yaml:"host"`     // Host para onde a solicitação será redirecionada
	Path       string           `yaml:"path"`
	StripPath  bool             `yaml:"stripPath"`
	Service    string           `yaml:"service"`
	TTL        int              `yaml:"ttl"`
	Port       string           `yaml:"port"`
	HealthPath string           `yaml:"healthPath"` // Caminho para healthcheck
	Retry      int              `yaml:"retry"`      // Número de tentativas para healthcheck
	RetryDelay int              `yaml:"retryDelay"` // Intervalo entre tentativas em segundos
	CORS       *cors.CORSConfig `yaml:"cors"`       // Configuração CORS para esta rota
}
