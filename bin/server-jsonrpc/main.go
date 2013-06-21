package main

import (
	"flag"
	"github.com/gorilla/handlers"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/client"
	"github.com/skynetservices/skynet/service"
	"log"
	"net/http"
	"os"
)

var (
	c        *client.Client
	flagset  = flag.NewFlagSet("main", flag.ContinueOnError)
	listen   = flagset.String("l", ":8080", "Address to listen on")
	s        = rpc.NewServer()
	services = make(map[string]*client.ServiceClient)
)

func init() {
	s.RegisterCodec(json.NewCodec(), "application/json")
}

func main() {
	StartClient()
	go StartService()

	flagsetArgs, skynetArgs := skynet.SplitFlagsetFromArgs(flagset, os.Args[1:])

	if err := flagset.Parse(flagsetArgs); err != nil {
		log.Fatal(err)
	}

	config, _ := skynet.GetClientConfigFromFlags(skynetArgs)

	config.Log = skynet.NewConsoleSemanticLogger("SkynetRPC", os.Stderr)
	c = client.NewClient(config)

	log.Print("Listening on " + *listen)
	http.Handle("/rpc", handlers.LoggingHandler(os.Stdout, s))
	log.Fatal(http.ListenAndServe(*listen, nil))
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
	c = client.NewClient(cConfig)
}

func StartService() {
	sConfig, _ := skynet.GetServiceConfig()
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
