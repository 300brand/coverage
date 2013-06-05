package main

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/client"
	"github.com/skynetservices/skynet/service"
	"os"
)

var (
	SCQueue *client.ServiceClient
	c       *client.Client
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
	config.Name = "Clock"
	config.Version = "1"

	s := &Clock{
		Log: skynet.NewConsoleSemanticLogger(config.Name, os.Stdout),
	}

	service := service.CreateService(s, config)
	defer service.Shutdown()

	waiter := service.Start(true)
	waiter.Wait()
}
