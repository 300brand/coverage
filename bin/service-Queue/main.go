package main

import (
	"git.300brand.com/coverage/config"
	"git.300brand.com/coverage/doozer/idqueue"
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
	c = client.NewClient(cConfig)
}

func StartService() {
	sConfig, _ := skynet.GetServiceConfig()
	sConfig.DoozerConfig.BootUri = config.Doozer.Address
	sConfig.Name = "Queue"
	sConfig.Version = "1"

	s := &Queue{
		Log: skynet.NewConsoleSemanticLogger(sConfig.Name, os.Stdout),
		FeedQ: &idqueue.IdQueue{
			Name: "feeds/" + sConfig.UUID,
			Addr: config.Doozer.Address,
			Max:  10,
		},
	}

	service := service.CreateService(s, sConfig)
	defer service.Shutdown()

	waiter := service.Start(true)
	waiter.Wait()
}
