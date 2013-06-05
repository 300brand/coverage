package main

import (
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
	config, _ := skynet.GetClientConfig()
	c = client.NewClient(config)
}

func StartService() {
	config, _ := skynet.GetServiceConfig()
	config.Name = "Queue"
	config.Version = "1"

	s := &Queue{
		Log: skynet.NewConsoleSemanticLogger(config.Name, os.Stdout),
		FeedQ: &idqueue.IdQueue{
			Name: "feeds",
			Addr: config.DoozerConfig.BootUri,
			Max:  10,
		},
	}

	service := service.CreateService(s, config)
	defer service.Shutdown()

	waiter := service.Start(true)
	waiter.Wait()
}
