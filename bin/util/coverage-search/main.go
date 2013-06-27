package main

import (
	"flag"
	"git.300brand.com/coverage/storage/mongo"
	"log"
)

var (
	dbHost, dbName string
	remap          bool
)

func init() {
	flag.StringVar(&dbHost, "dbHost", "localhost", "MongoDB host")
	flag.StringVar(&dbName, "dbName", "Coverage", "MongoDB database")
	flag.BoolVar(&remap, "remap", false, "Re-run Map-Reduce")
	log.SetFlags(log.Lmicroseconds)
}

func main() {
	flag.Parse()

	m := mongo.New(dbHost, dbName)
	if err := m.Connect(); err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	if remap {
		log.Print("Starting keyword map-reduce")
		info, err := m.ReduceKeywords(nil)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("InputCount:           %s", info.InputCount)
		log.Printf("EmitCount:            %s", info.EmitCount)
		log.Printf("OutputCount:          %s", info.OutputCount)
		log.Printf("VerboseTime.EmitLoop: %s", info.VerboseTime.EmitLoop)
		log.Printf("VerboseTime.Map:      %s", info.VerboseTime.Map)
		log.Printf("Time:                 %s", info.Time)
	}
}
