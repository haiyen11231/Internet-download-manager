package database

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/haiyen11231/Internet-download-manager/internal/utils"
	"go.uber.org/zap"
)

var (
	TabNameAccountPasswords = goqu.T("account_passwords")
)

const (
	ColNameAccountPasswordsOfAccountID  = "account_id"
	ColNameAccountPasswordsPasswordHash = "password_hash"
)

type AccountPassword struct {
	AccountID    uint64 `sql:"account_id"`
	PasswordHash string `sql:"password_hash"`
}

type AccountPasswordDataAccessor interface {
	CreateAccountPassword(ctx context.Context, accountPassword AccountPassword) error
	GetAccountPasssword(ctx context.Context, accountID uint64) (AccountPassword, error)
	UpdateAccountPassword(ctx context.Context, accountPassword AccountPassword) error
	WithDatabase(database Database) AccountPasswordDataAccessor
}

type accountPasswordDataAccessor struct {
	// database *goqu.Database
	database Database
	logger   *zap.Logger
}

func NewAccountPasswordDataAccessor(database *goqu.Database, logger *zap.Logger) AccountPasswordDataAccessor {
	return &accountPasswordDataAccessor{
		database: database,
		logger:   logger,
	}
}

func (a accountPasswordDataAccessor) CreateAccountPassword(ctx context.Context, accountPassword AccountPassword) error {
	logger := utils.LoggerWithContext(ctx, a.logger)
	_, err := a.database.Insert(TabNameAccountPasswords).Rows(goqu.Record{
		ColNameAccountPasswordsOfAccountID:  accountPassword.AccountID,
		ColNameAccountPasswordsPasswordHash: accountPassword.PasswordHash,
	}).Executor().ExecContext(ctx)

	if err != nil {
		logger.With(zap.Error(err)).Error("failed to create account password")
		return status.Errorf(codes.Internal, "failed to create account password: %+v", err)
	}

	return nil
}

// GetAccountPassword implements AccountPasswordDataAccessor.
func (a accountPasswordDataAccessor) GetAccountPassword(ctx context.Context, accountID uint64) (AccountPassword, error) {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Uint64("account_id", accountID))
	accountPassword := AccountPassword{}
	found, err := a.database.From(TabNameAccountPasswords).Where(goqu.Ex{ColNameAccountPasswordsOfAccountID: accountID}).ScanStructContext(ctx, &accountPassword)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get account password by account id")
		return AccountPassword{}, AccountPassword{}, status.Errorf(codes.Internal, "failed to get account password by account id: %+v", err)
	}

	if !found {
		logger.Warn("account password not found by account id")
		return AccountPassword{}, sql.ErrNoRows
	}

	return accountPassword, nil
}

func (a accountPasswordDataAccessor) UpdateAccountPassword(ctx context.Context, accountPassword AccountPassword) error {
	logger := utils.LoggerWithContext(ctx, a.logger)
	_, err := a.database.Update(TabNameAccountPasswords).Set(goqu.Record{
		ColNameAccountPasswordsPasswordHash: accountPassword.PasswordHash,
	}).Where(goqu.Ex{ColNameAccountPasswordsOfAccountID: accountPassword.AccountID}).Executor().ExecContext(ctx)

	if err != nil {
		logger.With(zap.Error(err)).Error("failed to update account password")
		return status.Errorf(codes.Internal, "failed to update account password: %+v", err)
	}

	return nil
}

func (a accountPasswordDataAccessor) WithDatabase(database Database) AccountPasswordDataAccessor {
	return &accountPasswordDataAccessor{
		database: database,
	}
}
