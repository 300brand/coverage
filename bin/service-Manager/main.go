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
	cConfig.DoozerConfig.BootUri = config.Doozer.Address
	cConfig.IdleConnectionsToInstance = 4
	cConfig.MaxConnectionsToInstance = cConfig.MaxConnectionsToInstance * 2
	c = client.NewClient(cConfig)
}

func StartService() {
	sConfig, _ := skynet.GetServiceConfig()
	sConfig.DoozerConfig.BootUri = config.Doozer.Address
	sConfig.Name = "Manager"
	sConfig.Version = "1"

	s := &Manager{
		Log: skynet.NewConsoleSemanticLogger(sConfig.Name, os.Stdout),
	}

	service := service.CreateService(s, sConfig)
	defer service.Shutdown()

	waiter := service.Start(true)
	waiter.Wait()
}
