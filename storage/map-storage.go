package storage

import (
	"sync"

	"github.com/palantir/stacktrace"

	api "github.com/gospodinbodurov/ports-apis/port-domain-service/api"
)

type MapStorage struct {
	storage map[string]*api.Port
	mutex   *sync.Mutex
}

func (ms *MapStorage) Init() error {
	ms.storage = map[string]*api.Port{}
	ms.mutex = &sync.Mutex{}

	return nil
}

func (ms *MapStorage) PutPort(port *api.Port) error {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	key := port.PortKey
	ms.storage[key] = port

	return nil
}

func (ms *MapStorage) GetPort(key string) (*api.Port, error) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	port, ok := ms.storage[key]

	if !ok {
		return nil, stacktrace.NewError("No such port")
	}

	return port, nil
}
