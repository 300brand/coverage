package main

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
	"log"
)

var services = make(map[string]func(*skynet.ServiceConfig) service.ServiceDelegate)

func main() {
	config, args := skynet.GetServiceConfig()
	log.Printf("%+v", args)
	config.Name = "StorageReader"

	if _, ok := services[config.Name]; !ok {
		log.Fatalf("Unknown service: %s", config.Name)
	}

	s := services[config.Name](config)
	srv := service.CreateService(s, config)

	defer func() {
		srv.Shutdown()
		if err := recover(); err != nil {
			log.Println("Unrecovered error occured: ", err)
		}
	}()

	waiter := srv.Start(true)
	waiter.Wait()
}
