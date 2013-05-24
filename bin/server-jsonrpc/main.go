package main

import (
	"flag"
	"github.com/gorilla/handlers"
	rpc "github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"log"
	"net/http"
	"os"
)

var (
	s      = rpc.NewServer()
	listen = flag.String("l", ":8080", "Address to listen on")
)

func init() {
	s.RegisterCodec(json.NewCodec(), "application/json")
}

func main() {
	flag.Parse()

	log.Print("Starting...")
	http.Handle("/rpc", handlers.LoggingHandler(os.Stdout, s))
	log.Fatal(http.ListenAndServe(*listen, nil))
}
