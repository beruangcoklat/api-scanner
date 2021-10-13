package repository

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func (r *apiDataRepository) PublishScanMessage(ctx context.Context, id string) error {
	err := r.kafkaWriter.WriteMessages(ctx, kafka.Message{
		Value: []byte(id),
	})
	if err != nil {
		return err
	}
	return nil
}
