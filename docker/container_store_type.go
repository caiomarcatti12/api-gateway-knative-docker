package docker

import (
	"sync"
)

type ContainerStore struct {
	containers map[string]Container // Mapeamento de ID do contêiner para detalhes
	mu         sync.RWMutex
}
