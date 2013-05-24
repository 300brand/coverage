package main

import (
	"flag"
	"github.com/gorilla/handlers"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
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
