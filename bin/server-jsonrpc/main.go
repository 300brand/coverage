package main

import (
	"flag"
	"github.com/gorilla/handlers"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/client"
	"log"
	"net/http"
	"os"
)

var (
	c        *client.Client
	config   *skynet.ClientConfig
	listen   = flag.String("l", ":8080", "Address to listen on")
	s        = rpc.NewServer()
	services = make(map[string]*client.ServiceClient)
)

func init() {
	s.RegisterCodec(json.NewCodec(), "application/json")
}

func main() {
	config, _ = skynet.GetClientConfig()

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
