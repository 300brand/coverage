package main

import (
	"git.300brand.com/coverage/node"
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
	"log"
	"os"
)

func main() {
	s := &node.SkynetFeed{}

	config, _ := skynet.GetServiceConfig()

	if config.Name == "" {
		config.Name = "SkynetFeed"
	}

	if config.Version == "unknown" {
		config.Version = "1"
	}

	if config.Region == "unknown" {
		config.Region = "Office"
	}

	var err error
	mlogger, err := skynet.NewMongoSemanticLogger("localhost", "skynet", "log", config.UUID)
	clogger := skynet.NewConsoleSemanticLogger(config.Name, os.Stdout)
	s.Log = skynet.NewMultiSemanticLogger(mlogger, clogger)
	config.Log = s.Log
	if err != nil {
		config.Log.Trace("Could not connect to mongo db for logging")
	}
	service := service.CreateService(s, config)

	// handle panic so that we remove ourselves from the pool in case
	// of catastrophic failure
	defer func() {
		service.Shutdown()
		if err := recover(); err != nil {
			log.Println("Unrecovered error occured: ", err)
		}
	}()

	// If we pass false here service will not be Registered we could
	// do other work/tasks by implementing the Started method and
	// calling Register() when we're ready
	waiter := service.Start(true)

	// waiting on the sync.WaitGroup returned by service.Start() will
	// wait for the service to finish running.
	waiter.Wait()
}
