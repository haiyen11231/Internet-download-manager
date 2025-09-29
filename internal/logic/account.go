package logic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/doug-martin/goqu/v9"
	"github.com/haiyen11231/Internet-download-manager/internal/data_access/cache"
	"github.com/haiyen11231/Internet-download-manager/internal/data_access/database"
	"github.com/haiyen11231/Internet-download-manager/internal/utils"
	"go.uber.org/zap"
)

// tao su tach biet giua layer grpc va logic
// Layer logic nay co the dc dung o layer grpc, kafka hoac cronjob
// khong muon viet logic bi phu thuoc boi logic toi tu dau
type CreateAccountParams struct {
	AccountName string
	Password    string
}

type CreateAccountOutput struct {
	ID        uint64
	AccountName string
}

type CreateSessionParams struct {
	AccountName string
	Password    string
}

type Account interface {
	CreateAccount(ctx context.Context, params CreateAccountParams) (CreateAccountOutput, error)
	CreateSession(ctx context.Context, params CreateSessionParams) (token string, err error)
}

type account struct {
	goquDatabase                *goqu.Database
	takenAccountNameCache       cache.TakenAccountName
	accountDataAccessor      database.AccountDataAccessor
	accountPasswordDataAccessor database.AccountPasswordDataAccessor
	hashLogic               Hash
	tokenLogic              Token
	logger                  *zap.Logger
}

// constructor
func NewAccount(goquDatabase *goqu.Database, takenAccountNameCache cache.TakenAccountName, accountDataAccessor database.AccountDataAccessor, accountPasswordDataAccessor database.AccountPasswordDataAccessor, hashLogic Hash, tokenLogic Token, logger *zap.Logger) Account {
	return &account{
		goquDatabase:             goquDatabase,
		takenAccountNameCache:    takenAccountNameCache,
		accountDataAccessor:      accountDataAccessor,
		accountPasswordDataAccessor: accountPasswordDataAccessor,
		hashLogic:               hashLogic,
		tokenLogic:              tokenLogic,
		logger:                  logger,
	}
}

func (a account) isAccountNameTaken (ctx context.Context, accountName string) (bool, error) {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.String("account_name", accountName))
	accountNameTaken, err := a.takenAccountNameCache.Has(ctx, accountName)
	if err != nil {
		logger.With(zap.Error(err)).Warn("failed to get account name from taken set in cache, will fall back to database")		// fallback to db check
	} else {
		return accountNameTaken, nil
	}

	_, err = a.accountDataAccessor.GetAccountByAccountName(ctx, accountName)
	if err != nil {
		// Error when account name not exist
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		// Error of connection
		return false, err
	}

	err = a.takenAccountNameCache.Add(ctx, accountName)
	if err != nil {
		logger.With(zap.Error(err)).Warn("failed to add account name to taken set in cache")
	}
	return true, nil
}	

// cam logic nay vao layer grpc handler
func (a account) CreateAccount(ctx context.Context, params CreateAccountParams) (CreateAccountOutput, error) {
	// flow: create new transaction -> ktra account name taken chua 
	// neu roi thi return alr taken
	// neu chua thi insert account vaof db van dung transaction -> hash password -> insert password vao db van dung transaction

	// ktra transaction co loi k
	accountNameTaken, err := a.isAccountNameTaken(ctx, params.AccountName)
	// tra ve err duy nhat thoi -> new co vde gif xay ra trong transaction thi se rollback lai
	if err != nil {
		return CreateAccountOutput{}, status.Errorf(codes.Internal, "failed to check if account name is taken")
	}

	if accountNameTaken {
		return CreateAccountOutput{}, status.Error(codes.AlreadyExists, "account name is already taken")
	}

	var accountID uint64
	// Create new transaction
	txErr := a.goquDatabase.WithTx(func(td *goqu.TxDatabase) error {
		// boi tx cung cai dat cac func cua interface db minh da dinh nghia -> co the thay the db cua accountDataAccessor bang transaction db ma k can phai viet lai CreateAccount
		accountID, err = a.accountDataAccessor.WithDatabase(td).CreateAccount(ctx, database.Account{
			AccountName: params.AccountName,
		})
		if err != nil {
			return err
		}

		// khong trace hash tuwf output ra input dc, nhung 2 input giong nhau se co hash giong nhau
		hashedPassword, hashErr := a.hashLogic.Hash(ctx, params.Password)
		if hashErr != nil {
			return hashErr
		}

		err = a.accountPasswordDataAccessor.WithDatabase(td).CreateAccountPassword(ctx, database.AccountPassword{
			AccountID: accountID,
			PasswordHash: hashedPassword,
		})
		if err != nil {
			return err
		}
		return nil
	})

	if txErr != nil {
		return CreateAccountOutput{}, txErr
	}

	return CreateAccountOutput{ID: accountID, AccountName: params.AccountName}, nil
	
}

func (a account) CreateSession(ctx context.Context, params CreateSessionParams) (string, error) {
	// get user by username
	// check password
	// generate token
	existingAccount, err := a.accountDataAccessor.GetAccountByAccountName(ctx, params.AccountName)
	// havent implement: loi xay ra do k tim thay trong db, loi connect to db
	if err != nil {
		return "", err
	}

	existingAccountPassword, err := a.accountPasswordDataAccessor.GetAccountPassword(ctx, existingAccount.AccountID)
	if err != nil {
		return "", err
	}

	isHashEqual, err := a.hashLogic.IsHashEqual(ctx, params.Password, existingAccountPassword.PasswordHash) 
	if err != nil {
		return "", err
	}

	if !isHashEqual {
		return "", status.Error(codes.Unauthenticated, "incorrect password")
	}

	token, _, err := a.tokenLogic.GetToken(ctx, existingAccount.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}