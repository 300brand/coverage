package config

import (
	"fmt"
	"io/ioutil"
	"launchpad.net/goyaml"
	"log"
	"os"
)

type config struct {
	Doozer    *doozer
	Mongo     *mongo
	Zookeeper *zookeeper
}

type doozer struct {
	Address string
}

type mongo struct {
	Database string
	Host     string
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
	Zookeeper = zookeeper{
		Address: "",
	}
	all = &config{
		Doozer:    &Doozer,
		Mongo:     &Mongo,
		Zookeeper: &Zookeeper,
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

	if err := goyaml.Unmarshal(b, all); err != nil {
		return fmt.Errorf("Error processing %s: %s", filename, err)
	}
	return
}

func readConfig() {
	switch err := ReadFromFile(ConfigFilename); true {
	case os.IsNotExist(err):
		fmt.Printf("Could not find configuration file %s. Please create with the following:\n", ConfigFilename)
		b, err := goyaml.Marshal(all)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
		os.Exit(1)
	case err != nil:
		log.Fatalf("Error processing configuration: %s", err)
	}
}
