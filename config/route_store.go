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
		routes: make(map[string]Route),
	}
}
func (rs *RouteStore) Add(route Route) {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	rs.routes[route.Service] = route
}

func (rs *RouteStore) Remove(routeService string) {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	delete(rs.routes, routeService)
}

func (rs *RouteStore) Get(routeService string) (Route, bool) {
	rs.mu.RLock()
	defer rs.mu.RUnlock()
	route, exists := rs.routes[routeService]
	return route, exists
}

func (rs *RouteStore) GetByPath(desiredPath string) (Route, bool) {
	rs.mu.RLock()
	defer rs.mu.RUnlock()

	for _, route := range rs.routes {
		if route.Path == desiredPath {
			return route, true
		}
	}

	return Route{}, false
}

func (rs *RouteStore) GetAll() map[string]Route {
	rs.mu.RLock()
	defer rs.mu.RUnlock()

	// Retornando uma cópia para evitar mutações externas
	copiedRoutes := make(map[string]Route, len(rs.routes))
	for path, route := range rs.routes {
		copiedRoutes[path] = route
	}
	return copiedRoutes
}

func (rs *RouteStore) MatchPrefix(prefixAccess string) bool {
	rs.mu.RLock()
	defer rs.mu.RUnlock()

	for _, route := range rs.GetAll() {
		if strings.HasPrefix(route.Path, prefixAccess) && prefixAccess != "/" {
			return true
		}
	}
	return false
}
