package main

import (
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
	config, _ := skynet.GetClientConfig()
	config.IdleConnectionsToInstance = 4
	config.MaxConnectionsToInstance = config.MaxConnectionsToInstance * 2
	c = client.NewClient(config)
}

func StartService() {
	config, _ := skynet.GetServiceConfig()
	config.Name = "Manager"
	config.Version = "1"

	s := &Manager{
		Log: skynet.NewConsoleSemanticLogger(config.Name, os.Stdout),
	}

	service := service.CreateService(s, config)
	defer service.Shutdown()

	waiter := service.Start(true)
	waiter.Wait()
}
