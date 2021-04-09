package storage

import api "github.com/gospodinbodurov/ports-apis/port-domain-service/api"

type Storage interface {
	Init() error
	PutPort(port *api.Port) error
	GetPort(key string) (*api.Port, error)
}
