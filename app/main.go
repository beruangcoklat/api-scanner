package main

import (
	"context"
	"flag"

	apidatarepo "github.com/beruangcoklat/api-scanner/api_data/repository"
	apidatausecase "github.com/beruangcoklat/api-scanner/api_data/usecase"
	"github.com/beruangcoklat/api-scanner/config"
	"github.com/beruangcoklat/api-scanner/constant"
	"github.com/beruangcoklat/api-scanner/domain"
	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	kafkaWriter *kafka.Writer
	redisClient *redis.Client

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

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

func initKafka() {
	kafkaWriter = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{config.GetConfig().KafkaBrokerAddr},
		Topic:   constant.Topic,
	})
}

func initRedis(ctx context.Context) error {
	var (
		err           error
		redisAddr     = config.GetConfig().RedisAddr
		redisPassword = config.GetConfig().RedisPassword
	)

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
	})

	_, err = redisClient.Ping(ctx).Result()
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
	apiDataRepo = apidatarepo.New(mongoDB, kafkaWriter, redisClient)
}

func initUsecase() {
	apiDataUc = apidatausecase.New(apiDataRepo)
}

func main() {
	var app string
	flag.StringVar(&app, "app", "http", "app type (http/kafka)")
	flag.Parse()

	switch app {
	case "http":
		runHTTP()
	case "kafka":
		runKafka()
	}
}
