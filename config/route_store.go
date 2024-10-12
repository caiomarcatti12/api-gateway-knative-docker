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
	"strings"
	"sync"
)

var routeStore *RouteStore
var routeStoreOnce sync.Once

func GetRouteStore() *RouteStore {
	routeStoreOnce.Do(func() {
		routeStore = newRouteStore()
	})
	return routeStore
}

// tornando privado para garantir que somente GetRouteStore() possa criar uma instância
func newRouteStore() *RouteStore {
	return &RouteStore{
		routesByContainerName: make(map[string]Route),
		routesByPrefix:        make(map[string]Route),
	}
}
func (rs *RouteStore) Add(route Route) {
	rs.routesByContainerName[route.ContainerName] = route
	rs.routesByPrefix[route.Path] = route
}

func (rs *RouteStore) Remove(routeService string) {
	delete(rs.routesByContainerName, routeService)
}

func (rs *RouteStore) Get(routeService string) (Route, bool) {
	route, exists := rs.routesByContainerName[routeService]
	return route, exists
}

func (rs *RouteStore) GetRouteByPath(path string) Route {
	prefixSlipted := strings.Split(path, "/")

	if len(prefixSlipted) > 1 {
		prefix := "/" + prefixSlipted[1]
		if route, exists := rs.routesByPrefix[prefix]; exists {
			return route
		}
	}
	return Route{}
}
