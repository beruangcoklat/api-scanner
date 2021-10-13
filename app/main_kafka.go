package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	apidatakafka "github.com/beruangcoklat/api-scanner/api_data/delivery/kafka"
)

func runKafka() {
	var (
		err error
		ctx = context.Background()
	)

	err = initConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = initMongo(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = initRedis(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = mongoClient.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		err = redisClient.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	initRepo()
	initUsecase()

	apidatakafka.NewAPIDataHandler(apiDataUc)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
