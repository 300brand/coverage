package main

import (
	"git.300brand.com/coverage/config"
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
	"os"
)

func main() {
	sConfig, _ := skynet.GetServiceConfig()
	sConfig.DoozerConfig.BootUri = config.Doozer.Address
	sConfig.Name = "StorageReader"
	sConfig.Version = "1"

	s := &StorageReader{
		Log:       skynet.NewConsoleSemanticLogger(sConfig.Name, os.Stdout),
		MongoHost: config.Mongo.Host,
		MongoDb:   config.Mongo.Database,
	}

	service := service.CreateService(s, sConfig)
	defer service.Shutdown()

	waiter := service.Start(true)
	waiter.Wait()
}
