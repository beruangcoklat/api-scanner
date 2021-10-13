package repository

import (
	"context"
)

func (r *apiDataRepository) SetScanRunning(ctx context.Context, id string) (bool, error) {
	cmd := r.redisClient.Incr(ctx, id)
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() == 1, nil
}

func (r *apiDataRepository) IsScanRunning(ctx context.Context, id string) (bool, error) {
	cmd := r.redisClient.Exists(ctx, id)
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() == 1, nil
}

func (r *apiDataRepository) FinishScan(ctx context.Context, id string) error {
	cmd := r.redisClient.Del(ctx, id)
	err := cmd.Err()
	if err != nil {
		return err
	}
	return nil
}
