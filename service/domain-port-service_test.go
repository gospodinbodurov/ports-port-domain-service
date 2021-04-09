package service

import (
	"errors"
	"testing"

	api "github.com/gospodinbodurov/ports-apis/port-domain-service/api"

	"github.com/gospodinbodurov/ports-port-domain-service/storage"
)

func TestGetPortHandler(t *testing.T) {
	storage := &storage.MapStorage{}

	err := storage.Init()

	if err != nil {
		t.Fatal(err)
	}

	s := &DomainPortService{
		Storage: storage,
	}

	getPortRequest := &api.GetPortRequest{
		PortKey: "testPort",
	}

	_, err = s.GetPort(nil, getPortRequest)

	if err == nil {
		t.Fatal(errors.New("Must return not found"))
	}

	port := &api.Port{
		PortKey: "testPort1",
		Name:    "test",
		City:    "testCity",
	}

	putPortRequest := &api.PutPortRequest{
		Port: port,
	}

	_, err = s.PutPort(nil, putPortRequest)

	if err != nil {
		t.Fatal(errors.New("Must put without error"))
	}

	getPortRequest1 := &api.GetPortRequest{
		PortKey: "testPort1",
	}

	response, err := s.GetPort(nil, getPortRequest1)

	if err != nil {
		t.Fatal(errors.New("Must return without error"))
	}

	if response.Port == nil {
		t.Fatal(errors.New("Must non nil port"))
	}

	portResponse := response.Port

	if portResponse.City != "testCity" {
		t.Fatal(errors.New("Must return testCity"))
	}
}

func TestPutPortHandler(t *testing.T) {
	storage := &storage.MapStorage{}

	err := storage.Init()

	if err != nil {
		t.Fatal(err)
	}

	s := &DomainPortService{
		Storage: storage,
	}

	port := &api.Port{
		PortKey: "testPort1",
		Name:    "test",
		City:    "testCity",
	}

	putPortRequest := &api.PutPortRequest{
		Port: port,
	}

	_, err = s.PutPort(nil, putPortRequest)

	if err != nil {
		t.Fatal(errors.New("Must put without error"))
	}

	getPortRequest1 := &api.GetPortRequest{
		PortKey: "testPort1",
	}

	response, err := s.GetPort(nil, getPortRequest1)

	if err != nil {
		t.Fatal(errors.New("Must return without error"))
	}

	if response.Port == nil {
		t.Fatal(errors.New("Must non nil port"))
	}

	portResponse := response.Port

	if portResponse.City != "testCity" {
		t.Fatal(errors.New("Must return testCity"))
	}
}
