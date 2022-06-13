package main

import (
	"log"
	"supreme/account"
	"supreme/configuration"
	"supreme/router"
)

func main() {
	envConfiguration := configuration.EnvConfiguration{}
	newAccount := account.Account{}
	HTTPRouter := router.HTTPRouter{
		Account: newAccount,
	}

	if err := envConfiguration.Load(); err != nil {
		log.Fatal(err)
	}

	if err := HTTPRouter.Route(envConfiguration.Listen); err != nil {
		log.Fatal(err)
	}
}
