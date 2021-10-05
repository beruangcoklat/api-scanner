package domain

import (
	"context"
	"time"
)

type (
	APIDataRepository interface {
		Create(ctx context.Context, data APIData) (string, error)
		AddScanResult(ctx context.Context, id string, data APIDataScanResult) error
		Get(ctx context.Context) ([]APIData, error)
		GetByID(ctx context.Context, id string) (*APIData, error)
		PublishScanMessage(ctx context.Context, id string) error
	}

	APIDataUsecase interface {
		Create(ctx context.Context, data APIData) error
		Get(ctx context.Context) ([]APIData, error)
		GetByID(ctx context.Context, id string) (*APIData, error)
		PublishScanMessage(ctx context.Context, id string) error
		Scan(ctx context.Context, id string) error
	}
)

type (
	APIData struct {
		ID         string              `json:"id" bson:"_id,omitempty"`
		Name       string              `json:"name" bson:"name"`
		Data       string              `json:"data" bson:"data"`
		DBMS       string              `json:"dbms" bson:"dbms"`
		ScanResult []APIDataScanResult `json:"scan_result" bson:"scan_result"`
	}

	APIDataScanResult struct {
		CreatedAt    time.Time `json:"created_at" bson:"created_at"`
		Log          string    `json:"log" bson:"log"`
		IsVulnerable bool      `json:"is_vulnerable" bson:"is_vulnerable"`
	}
)
