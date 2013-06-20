package main

import (
	"git.300brand.com/coverage/config"
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/client"
	"github.com/skynetservices/skynet/service"
	"os"
)

var c *client.Client

func main() {
	StartClient()
	StartService()
}

func StartClient() {
	cConfig, _ := skynet.GetClientConfig()
	c = client.NewClient(cConfig)
}

func StartService() {
	sConfig, _ := skynet.GetServiceConfig()
	sConfig.Name = "StorageWriter"
	sConfig.Version = "1"

	s := &StorageWriter{
		Log:       skynet.NewConsoleSemanticLogger(sConfig.Name, os.Stdout),
		MongoHost: config.Mongo.Host,
		MongoDb:   config.Mongo.Database,
	}

	service := service.CreateService(s, sConfig)
	defer service.Shutdown()

	waiter := service.Start(true)
	waiter.Wait()
}
