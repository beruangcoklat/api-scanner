package usecase

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/beruangcoklat/api-scanner/config"
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

func (uc *apiDataUsecase) Get(ctx context.Context) ([]domain.APIData, error) {
	data, err := uc.apiDataRepo.Get(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *apiDataUsecase) PublishScanMessage(ctx context.Context, id string) error {
	success, err := uc.apiDataRepo.SetScanRunning(ctx, id)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("scan already running")
	}

	err = uc.apiDataRepo.PublishScanMessage(ctx, id)
	if err != nil {
		uc.apiDataRepo.FinishScan(ctx, id)
		return err
	}
	return nil
}

func (uc *apiDataUsecase) Scan(ctx context.Context, id string) error {
	apiData, err := uc.apiDataRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	filepath := fmt.Sprintf("/tmp/%v", time.Now().Unix())
	err = os.WriteFile(filepath, []byte(apiData.Data), 0644)
	if err != nil {
		return err
	}

	defer func() {
		os.Remove(filepath)
		uc.apiDataRepo.FinishScan(ctx, id)
	}()

	command := fmt.Sprintf("python3 sqlmap.py -r %v --dbms=%v --level=5 --risk=3 --flush-session --batch", filepath, apiData.DBMS)
	cmd := exec.CommandContext(ctx, "/bin/bash", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Dir = config.GetConfig().SqlmapPath

	var outb bytes.Buffer
	cmd.Stdout = &outb

	err = cmd.Run()
	if err != nil {
		return err
	}

	log := outb.String()
	err = uc.apiDataRepo.AddScanResult(ctx, id, domain.APIDataScanResult{
		CreatedAt:    time.Now(),
		Log:          log,
		IsVulnerable: strings.Contains(log, "---"),
	})
	if err != nil {
		return err
	}

	return nil
}

func (uc *apiDataUsecase) IsScanRunning(ctx context.Context, id string) (bool, error) {
	isRunning, err := uc.apiDataRepo.IsScanRunning(ctx, id)
	if err != nil {
		return false, err
	}
	return isRunning, nil
}
