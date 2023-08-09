package docker

import "time"

type Container struct {
	ID         string
	Service    string
	LastAccess time.Time
}
