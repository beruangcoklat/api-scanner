package repository

import (
	"context"

	"github.com/beruangcoklat/api-scanner/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *apiDataRepository) Create(ctx context.Context, data domain.APIData) (string, error) {
	result, err := r.mongoDB.Collection("api_data").InsertOne(ctx, data)
	if err != nil {
		return "", err
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", nil
	}

	return objectID.Hex(), nil
}

func (r *apiDataRepository) AddScanResult(ctx context.Context, id string, data domain.APIDataScanResult) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.mongoDB.Collection("api_data").UpdateByID(ctx, objectID, bson.M{
		"$push": bson.M{
			"scan_result": data,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *apiDataRepository) Get(ctx context.Context) ([]domain.APIData, error) {
	cursor, err := r.mongoDB.Collection("api_data").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	result := []domain.APIData{}
	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *apiDataRepository) GetByID(ctx context.Context, id string) (*domain.APIData, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := r.mongoDB.Collection("api_data").FindOne(ctx, bson.M{
		"_id": objectID,
	})
	if err != nil {
		return nil, err
	}

	var data *domain.APIData
	err = result.Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
