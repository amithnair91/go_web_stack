package storage_test

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
)

type MongoTestContainer struct {
	IP      string
	Port    string
	mongo   testcontainers.Container
	context context.Context
}

func (mtc *MongoTestContainer) Start(t *testing.T) (mongo testcontainers.Container, ip string, port string) {
	port = "27017"
	ctx := context.Background()
	mtc.context = ctx
	req := testcontainers.ContainerRequest{
		Image:        "mongo:4.2.6-bionic",
		ExposedPorts: []string{fmt.Sprintf("%s/tcp", port)},
		WaitingFor:   wait.ForListeningPort(nat.Port(port)),
	}
	mongo, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	mtc.mongo = mongo
	if err != nil {
		t.Error(err)
	}
	ip, err = mongo.Host(ctx)
	if err != nil {
		t.Error(err)
	}
	_, err = mongo.MappedPort(ctx, nat.Port(port))
	if err != nil {
		t.Error(err)
	}
	return
}

func (mtc *MongoTestContainer) Stop() {
	mtc.mongo.Terminate(mtc.context)
}
