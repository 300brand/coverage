package main

import (
	"git.300brand.com/coverage/node"
	"git.300brand.com/coverage/storage/mongo"
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/service"
	"log"
	"os"
)

func init() {
	services["StorageReader"] = storageReader
}

func storageReader(config *skynet.ServiceConfig) service.ServiceDelegate {
	sr := &node.StorageReader{
		Log:   skynet.NewConsoleSemanticLogger(config.Name, os.Stdout),
		Mongo: mongo.New("localhost", "Coverage"),
	}
	if err := sr.Mongo.Connect(); err != nil {
		log.Fatal(err)
	}
	return sr
}
