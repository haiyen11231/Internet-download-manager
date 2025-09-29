package cache

import (
	"context"

	"github.com/haiyen11231/Internet-download-manager/internal/utils"
	"go.uber.org/zap"
)

const setKeyNameTakenAccountName = "taken_account_name_set"

type TakenAccountName interface {
	Add(ctx context.Context, accountName string) error
	Has(ctx context.Context, accountName string) (bool, error)
}

type takenAccountName struct {
	client Client
	logger *zap.Logger
}

func NewTakenAccountName(client Client, logger *zap.Logger) TakenAccountName {
	return &takenAccountName{
		client: client,
		logger: logger,
	}
}

func (t takenAccountName) Add(ctx context.Context, accountName string) error {
	logger := utils.LoggerWithContext(ctx, t.logger).With(zap.String("account_name", accountName))

	if err := t.client.AddToSet(ctx, setKeyNameTakenAccountName, accountName); err != nil {
		logger.With(zap.Error(err)).Error("failed to add account name to set in cache")
		return err
	}

	return nil
}

func (t takenAccountName) Has(ctx context.Context, accountName string) (bool, error) {
	logger := utils.LoggerWithContext(ctx, t.logger).With(zap.String("account_name", accountName))
	
	result, err := t.client.IsDataInSet(ctx, setKeyNameTakenAccountName, accountName)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to check if account name is in set in cache")
		return false, err
	}

	return result, nil
}