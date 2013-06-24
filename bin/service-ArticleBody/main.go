package main

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
	"log"
	"os"
)

func main() {
	s := &ArticleBody{}
	config, _ := skynet.GetServiceConfig()
	config.Name = "ArticleBody"
	config.Version = "1"

	s.Log = skynet.NewConsoleSemanticLogger(config.Name, os.Stdout)
	config.Log = s.Log
	service := service.CreateService(s, config)

	defer func() {
		service.Shutdown()
		if err := recover(); err != nil {
			log.Println("Unrecovered error occured: ", err)
		}
	}()

	waiter := service.Start(true)
	waiter.Wait()
}
