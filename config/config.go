package config

import (
	"fmt"
	"io/ioutil"
	"launchpad.net/goyaml"
	"log"
	"os"
	"time"
)

type doozer struct {
	Address string
}

type mongo struct {
	Database string
	Host     string
}

type rpcserver struct {
	Address string
}

type timeouts struct {
	Download time.Duration
}

type zookeeper struct {
	Address string
}

const ConfigFilename = "/etc/coverage.yaml"

var (
	Doozer = doozer{
		Address: "localhost:8046",
	}
	Mongo = mongo{
		Database: "Coverage",
		Host:     "localhost",
	}
	RPCServer = rpcserver{
		Address: ":8080",
	}
	Timeouts = timeouts{
		Download: time.Minute,
	}
	Zookeeper = zookeeper{
		Address: "",
	}
	config = &struct {
		Doozer    *doozer
		Mongo     *mongo
		RPCServer *rpcserver
		Timeouts  *timeouts
		Zookeeper *zookeeper
	}{
		&Doozer,
		&Mongo,
		&RPCServer,
		&Timeouts,
		&Zookeeper,
	}
)

func init() {
	readConfig()
}

func ReadFromFile(filename string) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("Error reading %s: %s", filename, err)
	}

	if err := goyaml.Unmarshal(b, config); err != nil {
		return fmt.Errorf("Error processing %s: %s", filename, err)
	}
	return
}

func readConfig() {
	switch err := ReadFromFile(ConfigFilename); true {
	case os.IsNotExist(err):
		fmt.Printf("Could not find configuration file %s. Please create with the following:\n", ConfigFilename)
		b, err := goyaml.Marshal(config)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
		os.Exit(1)
	case err != nil:
		log.Fatalf("Error processing configuration: %s", err)
	}
}
