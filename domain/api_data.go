package domain

import (
	"context"
	"time"
)

type (
	APIDataRepository interface {
		Create(ctx context.Context, data APIData) (string, error)
		AddTestResult(ctx context.Context, id string, data APIDataTestResult) error
		Get(ctx context.Context) ([]APIData, error)
		GetByID(ctx context.Context, id string) (*APIData, error)
	}

	APIDataUsecase interface {
		Create(ctx context.Context, data APIData) error
		AddTestResult(ctx context.Context, id string, data APIDataTestResult) error
		Get(ctx context.Context) ([]APIData, error)
		GetByID(ctx context.Context, id string) (*APIData, error)
	}
)

type (
	APIData struct {
		ID         string              `json:"id" bson:"_id,omitempty"`
		Name       string              `json:"name" bson:"name"`
		Data       string              `json:"data" bson:"data"`
		TestResult []APIDataTestResult `json:"test_result" bson:"test_result"`
	}

	APIDataTestResult struct {
		CreatedAt    time.Time `json:"created_at" bson:"created_at"`
		Log          string    `json:"log" bson:"log"`
		IsVulnerable bool      `json:"is_vulnerable" bson:"is_vulnerable"`
	}
)
