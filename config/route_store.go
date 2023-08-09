package config

import "sync"

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
