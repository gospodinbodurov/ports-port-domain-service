package storage

import (
	"github.com/palantir/stacktrace"

	api "github.com/gospodinbodurov/ports-apis/port-domain-service/api"
)

type MapStorage struct {
	storage map[string]*api.Port
}

func (ms *MapStorage) Init() error {
	ms.storage = map[string]*api.Port{}

	return nil
}

func (ms *MapStorage) PutPort(port *api.Port) error {
	key := port.PortKey
	ms.storage[key] = port

	return nil
}

func (ms *MapStorage) GetPort(key string) (*api.Port, error) {
	port, ok := ms.storage[key]

	if !ok {
		return nil, stacktrace.NewError("No such port")
	}

	return port, nil
}
