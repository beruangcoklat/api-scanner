package usecase

import (
	"context"

	"github.com/beruangcoklat/api-scanner/domain"
)

type apiDataUsecase struct {
	apiDataRepo domain.APIDataRepository
}

func New(apiDataRepo domain.APIDataRepository) domain.APIDataUsecase {
	return &apiDataUsecase{
		apiDataRepo: apiDataRepo,
	}
}

func (uc *apiDataUsecase) Create(ctx context.Context, data domain.APIData) error {
	_, err := uc.apiDataRepo.Create(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (uc *apiDataUsecase) GetByID(ctx context.Context, id string) (*domain.APIData, error) {
	data, err := uc.apiDataRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *apiDataUsecase) AddTestResult(ctx context.Context, id string, data domain.APIDataTestResult) error {
	err := uc.apiDataRepo.AddTestResult(ctx, id, data)
	if err != nil {
		return err
	}
	return nil
}

func (uc *apiDataUsecase) Get(ctx context.Context) ([]domain.APIData, error) {
	data, err := uc.apiDataRepo.Get(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}
