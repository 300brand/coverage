package main

import (
	"fmt"
	"git.300brand.com/coverage"
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/skynet/client"
	"labix.org/v2/mgo/bson"
	"os"
)

func main() {
	config, _ := skynet.GetClientConfig()

	config.Log = skynet.NewConsoleSemanticLogger("TestServiceClient", os.Stderr)

	client := client.NewClient(config)

	// This will not fail if no services currently exist, as
	// connections are created on demand this saves from chicken and
	// egg issues with dependencies between services
	service := client.GetService("SkynetFeed", "", "", "")
	// (any version, any region, any host)

	in := map[bson.ObjectId]bool{}
	out := &coverage.Feed{}

	if err := service.Send(nil, "NextFeed", in, out); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", out)

	if err := service.Send(nil, "Download", out, out); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", out)
}
