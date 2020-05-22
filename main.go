package main

import (
	"fmt"

	"github.com/Zero23ku/estructura-test/pkg/config"
	"github.com/Zero23ku/estructura-test/pkg/server"
)

func main() {
	conf, err := config.LoadConfig("./config.json")
	if err != nil {
		fmt.Println("Error!")
	} else {
		server := &server.Server{}
		server.Initialize()
		server.Run(conf.Port)
	}

}
