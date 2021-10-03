package http

import (
	"context"
	"log"
	"os"

	"github.com/beruangcoklat/api-scanner/config"
	"github.com/beruangcoklat/api-scanner/constanta"
	"github.com/beruangcoklat/api-scanner/domain"
	"github.com/segmentio/kafka-go"
)

type apiDataHandler struct {
	apiDataUsecase domain.APIDataUsecase
}

func NewAPIDataHandler(apiDataUsecase domain.APIDataUsecase) {
	handler := &apiDataHandler{
		apiDataUsecase: apiDataUsecase,
	}

	ctx := context.Background()
	go handler.SubscribeScanMessage(ctx)
}

func (h *apiDataHandler) SubscribeScanMessage(ctx context.Context) {
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{config.GetConfig().KafkaBrokerAddr},
		Topic:   constanta.Topic,
		GroupID: "my-group",
		Logger:  log.New(os.Stdout, "kafka reader: ", 0),
	})

	for {
		msg, err := kafkaReader.ReadMessage(ctx)
		if err != nil {
			log.Fatal(err)
		}

		id := string(msg.Value)
		err = h.apiDataUsecase.Scan(ctx, id)
		if err != nil {
			log.Fatal(err)
		}
	}
}
