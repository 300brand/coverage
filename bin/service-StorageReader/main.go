package main

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
	"os"
)

var (
	mongoDb   = skynet.GetDefaultEnvVar("SERVICE_MONGO_DB", "Coverage")
	mongoHost = skynet.GetDefaultEnvVar("SERVICE_MONGO_HOST", "localhost")
)

func main() {
	config, _ := skynet.GetServiceConfig()
	config.Name = "StorageReader"
	config.Version = "1"

	s := &StorageReader{
		Log:       skynet.NewConsoleSemanticLogger(config.Name, os.Stdout),
		MongoHost: mongoHost,
		MongoDb:   mongoDb,
	}

	service := service.CreateService(s, config)
	defer service.Shutdown()

	waiter := service.Start(true)
	waiter.Wait()
}
