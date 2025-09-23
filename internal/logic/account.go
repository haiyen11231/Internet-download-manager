package logic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/doug-martin/goqu/v9"
	"github.com/haiyen11231/Internet-download-manager/internal/data_access/database"
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
	accountDataAccessor      database.AccountDataAccessor
	accountPasswordDataAccessor database.AccountPasswordDataAccessor
	hashLogic               Hash
	tokenLogic              Token
}

// constructor
func NewAccount(goquDatabase *goqu.Database, accountDataAccessor database.AccountDataAccessor, accountPasswordDataAccessor database.AccountPasswordDataAccessor, hashLogic Hash, tokenLogic Token) Account {
	return &account{
		goquDatabase:             goquDatabase,
		accountDataAccessor:      accountDataAccessor,
		accountPasswordDataAccessor: accountPasswordDataAccessor,
		hashLogic:               hashLogic,
		tokenLogic:              tokenLogic,
	}
}

func (a account) isAccountNameTaken (ctx context.Context, accountName string) (bool, error) {
	_, err := a.accountDataAccessor.GetAccountByAccountName(ctx, accountName)
	if err != nil {
		// Error when account name not exist
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		// Error of connection
		return false, err
	}
	return true, nil
}	

// cam logic nay vao layer grpc handler
func (a account) CreateAccount(ctx context.Context, params CreateAccountParams) (CreateAccountOutput, error) {
	// flow: create new transaction -> ktra account name taken chua 
	// neu roi thi return alr taken
	// neu chua thi insert account vaof db van dung transaction -> hash password -> insert password vao db van dung transaction

	// ktra transaction co loi k
	var accountID uint64
	// Create new transaction
	txErr := a.goquDatabase.WithTx(func(td *goqu.TxDatabase) error {
		isTaken, err := a.isAccountNameTaken(ctx, params.AccountName);
		// tra ve err duy nhat thoi -> new co vde gif xay ra trong transaction thi se rollback lai
		if err != nil {
			return err
		} 
		
		if isTaken {
			return errors.New("account name is already taken")
		}

		// boi tx cung cai dat cac func cua interface db minh da dinh nghia -> co the thay the db cua accountDataAccessor bang transaction db ma k can phai viet lai CreateAccount
		accountID, err = a.accountDataAccessor.WithDatabase(td).CreateAccount(ctx, database.Account{
			AccountName: params.AccountName,
		})
		if err != nil {
			return err
		}

		// khong trace hash tuwf output ra input dc, nhung 2 input giong nhau se co hash giong nhau
		hashedPassword, err := a.hashLogic.Hash(ctx, params.Password)
		if err != nil {
			return err
		}

		if err := a.accountPasswordDataAccessor.WithDatabase(td).CreateAccountPassword(ctx, database.AccountPassword{
			AccountID: accountID,
			PasswordHash: hashedPassword,
		}); err != nil {
			return err
		}
		return nil
	})

	if txErr != nil {
		return CreateAccountOutput{}, txErr
	}

	return CreateAccountOutput{ID: accountID, AccountName: params.AccountName}, nil
	
}

func (a account) CreateSession(ctx context.Context, params CreateSessionParams) (token string, err error) {
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
		return "", errors.New("invalid password")
	}

	return "", nil
}