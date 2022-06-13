package main

import (
	"log"
	"supreme/configuration"
)

func main() {
	newConfiguration := configuration.EnvConfiguration{}

	if err := newConfiguration.Load(); err != nil {
		log.Fatal(err)
	}
}
