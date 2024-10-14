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

import (
	"fmt"
	"strings"
	"sync"
)

// HostStore é o armazenamento principal para hosts e rotas.
type HostStore struct {
	store map[string]HostData
}

// HostData armazena as rotas e a configuração de CORS para cada host.
type HostData struct {
	CORS   CORSConfig             // Configuração CORS específica do host
	Routes map[string]RouteConfig // Mapeamento de rotas pelo path
}

var (
	once     sync.Once
	instance *HostStore
)

// GetHostStore retorna a instância Singleton de HostStore.
func GetHostStore() *HostStore {
	once.Do(func() {
		instance = &HostStore{
			store: make(map[string]HostData),
		}
	})
	return instance
}

// AddHost adiciona ou atualiza um host no HostStore.
func (hs *HostStore) AddHost(hostConfig HostConfig) {
	// Criar o map de rotas para o host.
	routeMap := make(map[string]RouteConfig)
	for _, route := range hostConfig.Routes {
		routeMap[route.Path] = route
	}

	// Adicionar o host com suas rotas e CORS.
	hs.store[hostConfig.Host] = HostData{
		CORS:   hostConfig.CORS,
		Routes: routeMap,
	}
}

// GetRoute obtém uma rota específica de um host pelo path.
func (hs *HostStore) GetRoute(host, path string) (RouteConfig, bool) {
	hostData, ok := hs.store[host]
	if !ok {
		return RouteConfig{}, false
	}

	prefix := hs.getPrefixPath(path)

	route, found := hostData.Routes[prefix]

	return route, found
}

// GetAllRoutes obtém todas as rotas de um host específico.
func (hs *HostStore) GetAllRoutes(host string) ([]RouteConfig, bool) {
	hostData, ok := hs.store[host]
	if !ok {
		return nil, false
	}

	// Converter o map de rotas para uma slice.
	routes := make([]RouteConfig, 0, len(hostData.Routes))
	for _, route := range hostData.Routes {
		routes = append(routes, route)
	}
	return routes, true
}

// GetCORS obtém a configuração de CORS de um host.
func (hs *HostStore) GetCORS(host string) (CORSConfig, bool) {
	hostData, ok := hs.store[host]
	if !ok {
		return CORSConfig{}, false
	}
	return hostData.CORS, true
}

// ListHosts retorna todos os hosts armazenados.
func (hs *HostStore) ListHosts() []string {
	hosts := make([]string, 0, len(hs.store))
	for host := range hs.store {
		hosts = append(hosts, host)
	}
	return hosts
}

func (hs *HostStore) getPrefixPath(path string) string {
	prefixSlipted := strings.Split(path, "/")

	if len(prefixSlipted) > 1 {
		return fmt.Sprintf("/%s", prefixSlipted[1])

	}
	return ""
}
