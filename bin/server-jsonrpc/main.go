package main

import (
	"git.300brand.com/coverage/config"
	"github.com/gorilla/handlers"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/client"
	"github.com/skynetservices/skynet/service"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	c        *client.Client
	s        = rpc.NewServer()
	services = make(map[string]*client.ServiceClient)
)

func init() {
	s.RegisterCodec(json.NewCodec(), "application/json")
}

func main() {
	listener, err := net.Listen("tcp", config.RPCServer.Address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	go func(l net.Listener) {
		log.Printf("Listening on %s", l)
		http.Handle("/rpc", handlers.LoggingHandler(os.Stdout, s))
		log.Fatal(http.Serve(l, nil))
	}(listener)

	StartClient()
	StartService()
}

func GetService(name string) (s *client.ServiceClient) {
	s, ok := services[name]
	if !ok {
		s = c.GetService(name, "", "", "")
		services[name] = s
	}
	return
}

func StartClient() {
	cConfig, _ := skynet.GetClientConfig()
	cConfig.DoozerConfig.BootUri = config.Doozer.Address
	c = client.NewClient(cConfig)
}

func StartService() {
	sConfig, _ := skynet.GetServiceConfig()
	sConfig.DoozerConfig.BootUri = config.Doozer.Address
	sConfig.Name = "ServerJSONRPC"
	sConfig.Version = "1"

	s := &ServerJSONRPC{
		Log: skynet.NewConsoleSemanticLogger(sConfig.Name, os.Stdout),
	}

	service := service.CreateService(s, sConfig)
	defer service.Shutdown()

	waiter := service.Start(true)
	waiter.Wait()
}
