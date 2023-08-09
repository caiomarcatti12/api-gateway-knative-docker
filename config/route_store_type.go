package config

import "sync"

type RouteStore struct {
	routes map[string]Route // Mapeamento do caminho da rota para detalhes
	mu     sync.RWMutex
}
