package main

import (
	"github.com/Zero23ku/estructura-test/pkg/server"
)

func main() {
	server := &server.Server{}
	server.Initialize()
	server.Run(":8080")

}
