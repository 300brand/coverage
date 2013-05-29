package main

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/client"
	"github.com/skynetservices/skynet/service"
	"os"
)

var (
	c         *client.Client
	mongoDb   = skynet.GetDefaultEnvVar("SERVICE_MONGO_DB", "Coverage")
	mongoHost = skynet.GetDefaultEnvVar("SERVICE_MONGO_HOST", "localhost")
)

func main() {
	StartClient()
	StartService()
}

func StartClient() {
	config, _ := skynet.GetClientConfig()
	c = client.NewClient(config)
}

func StartService() {
	config, _ := skynet.GetServiceConfig()
	config.Name = "StorageWriter"
	config.Version = "1"

	s := &StorageWriter{
		Log:       skynet.NewConsoleSemanticLogger(config.Name, os.Stdout),
		MongoHost: mongoHost,
		MongoDb:   mongoDb,
	}

	service := service.CreateService(s, config)
	defer service.Shutdown()

	waiter := service.Start(true)
	waiter.Wait()
}
