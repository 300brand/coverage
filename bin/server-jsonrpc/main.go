package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
	"os"
)

var s = rpc.NewServer()

func init() {
	s.RegisterCodec(json.NewCodec(), "application/json")
}

func main() {
	log.Print("Starting...")
	http.Handle("/rpc", handlers.LoggingHandler(os.Stdout, s))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
