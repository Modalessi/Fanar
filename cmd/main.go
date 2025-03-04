package main

import (
	fanar "github.com/Modalessi/iau_resources/fanar_api"
)

func main() {
	fanarServer := fanar.NewFanarServer(":4000")

	err := fanarServer.Start()
	if err != nil {
		panic("error starting fanar server")
	}
}
