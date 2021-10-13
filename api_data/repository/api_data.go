package repository

import (
	"github.com/beruangcoklat/api-scanner/domain"
	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type apiDataRepository struct {
	mongoDB     *mongo.Database
	kafkaWriter *kafka.Writer
	redisClient *redis.Client
}

func New(
	mongoDB *mongo.Database,
	kafkaWriter *kafka.Writer,
	redisClient *redis.Client,
) domain.APIDataRepository {

	return &apiDataRepository{
		mongoDB:     mongoDB,
		kafkaWriter: kafkaWriter,
		redisClient: redisClient,
	}
}
