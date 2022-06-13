package main

import (
	"log"
	"supreme/account"
	"supreme/configuration"
	"supreme/router"
	// redisClient "supreme/redis"
	// "github.com/go-redis/redis"
)

func main() {
	// configuration
	envConfiguration := configuration.EnvConfiguration{}

	// external libraries
	// newRedisClient := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "P@ssw0rd",
	// 	DB:       0,
	// })

	// // project packages
	// redisStorage := redisClient.RedisStorage{
	// 	Client: newRedisClient,
	// }
	// middleware to make read and write
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
