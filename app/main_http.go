package main

import (
	"context"
	"log"
	"net/http"

	apidatahttp "github.com/beruangcoklat/api-scanner/api_data/delivery/http"
	"github.com/beruangcoklat/api-scanner/config"
	"github.com/gorilla/mux"
)

func runHTTP() {
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

	initKafka()

	defer func() {
		err = mongoClient.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		err = kafkaWriter.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	initRepo()
	initUsecase()

	port := config.GetConfig().Port
	router := mux.NewRouter()
	apidatahttp.NewAPIDataHandler(router, apiDataUc)

	log.Print("listen :" + port)
	http.ListenAndServe(":"+port, router)
}
