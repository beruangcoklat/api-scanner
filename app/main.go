package main

import (
	"context"
	"log"
	"net/http"

	apidatahandler "github.com/beruangcoklat/api-scanner/api_data/delivery/http"
	apidatarepo "github.com/beruangcoklat/api-scanner/api_data/repository"
	apidatausecase "github.com/beruangcoklat/api-scanner/api_data/usecase"
	"github.com/beruangcoklat/api-scanner/config"
	"github.com/beruangcoklat/api-scanner/domain"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client

	apiDataRepo domain.APIDataRepository
	apiDataUc   domain.APIDataUsecase
)

func initConfig() error {
	return config.Init("/etc/api_scanner/config.json")
}

func initMongo(ctx context.Context) error {
	var (
		err      error
		mongoURI = config.GetConfig().MongoURI
	)

	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}

	return nil
}

func initRepo() {
	var (
		mongoDBName = config.GetConfig().MongoDbName
	)

	mongoDB := mongoClient.Database(mongoDBName)
	apiDataRepo = apidatarepo.New(mongoDB)
}

func initUsecase() {
	apiDataUc = apidatausecase.New(apiDataRepo)
}

func initHandler() {
	var (
		port = config.GetConfig().Port
	)

	router := mux.NewRouter()
	apidatahandler.NewAPIDataHandler(router, apiDataUc)

	log.Print("listen :" + port)
	http.ListenAndServe(":"+port, router)
}

func main() {
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

	defer func() {
		err = mongoClient.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	initRepo()
	initUsecase()
	initHandler()
}
