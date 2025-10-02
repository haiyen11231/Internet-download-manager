package database

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/haiyen11231/Internet-download-manager/internal/utils"
	"go.uber.org/zap"
)

var (
	TabNameAccounts = goqu.T("accounts")
	ErrAccountNotFound = status.Error(codes.NotFound, "account not found")
)

const (
	ColNameAccountsID 	 = "account_id"
	ColNameAccountsAccountName    = "account_name"
)

type Account struct {
	AccountID   uint64 `db:"account_id" goqu:"skipinsert,skipupdate"`
	AccountName string `db:"account_name"`
}

type AccountDataAccessor interface {
	CreateAccount(ctx context.Context, account Account) (uint64, error)
	GetAccountByID(ctx context.Context, accountID uint64) (Account, error)
	GetAccountByAccountName(ctx context.Context, accountName string) (Account, error)
	// cai dat transaction - Query password that bai -> cac queries khacs phai revert lai -> cac queris cung thanh cong/cung that bai
	WithDatabase(database Database) AccountDataAccessor
	// transaction db va db bthg se chung 1 so command voi nhau
	// muon dung chung cau lenh GetAccountByAccountName hay GetAccountByID trong db goqu bthg hay transaction db thi van muon dung chung cau lenh nhu the
	// -> can interface chung de tru func chung giuwa db bthg va db transaction
}

type accountDataAccessor struct {
	database Database // co th tra ve transaction database -> interact with transaction, thay vi tuong tac voi goqu thi tuong tac voi db bthg
	logger *zap.Logger
}

func NewAccountDataAccessor(database *goqu.Database, logger *zap.Logger) AccountDataAccessor {
	return &accountDataAccessor{
		database: database,
		logger:   logger,
	}
}

func (a accountDataAccessor) CreateAccount(ctx context.Context, account Account) (uint64, error) {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("account", account))

	result, err := a.database.Insert(TabNameAccounts).Rows(goqu.Record{
		ColNameAccountsAccountName: account.AccountName,
	}).Executor().ExecContext(ctx)

	if err != nil {
		logger.With(zap.Error(err)).Error("failed to create account")
		return 0, status.Error(codes.Internal, "failed to create account")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get last insert id after creating account")
		return 0, status.Error(codes.Internal, "failed to get last insert id")
	}

	return uint64(id), nil
}

func (a accountDataAccessor) GetAccountByID(ctx context.Context, accountID uint64) (Account, error) {
	// implement get account by id in db
	logger := utils.LoggerWithContext(ctx, a.logger)
	account := Account{}
	found, err := a.database.From(TabNameAccounts).Where(goqu.C(ColNameAccountsID).Eq(accountID)).ScanStructContext(ctx, &account)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get account by id")
		return Account{}, status.Error(codes.Internal, "failed to get account by id")
	}

	if !found {
		logger.Warn("account not found by id")
		return Account{}, ErrAccountNotFound
	}

	return account, nil
}		

func (a accountDataAccessor) GetAccountByAccountName(ctx context.Context, accountName string) (Account, error) {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.String("account_name", accountName))
	account := Account{}
	found, err := a.database.From(TabNameAccounts).Where(goqu.C(ColNameAccountsAccountName).Eq(accountName)).ScanStructContext(ctx, &account)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get account by account name")
		return Account{}, status.Error(codes.Internal, "failed to get account by account name")
	}

	if !found {
		logger.Warn("account not found by account name")
		return Account{}, ErrAccountNotFound
	}
	// implement get account by username in db
	return Account{}, nil
}

// why? khi thay the db cua minh bang transaction db thi van nhan dc doi tuong AccountDataAccessor -> nhung func con lai van dc cai dat dua tren database interface cua minh thay vi dua tren 
// datacbase struct cua goqu
func (a accountDataAccessor) WithDatabase(database Database) AccountDataAccessor {
	return &accountDataAccessor{
		database: database,
		logger:   a.logger,
	}
}
