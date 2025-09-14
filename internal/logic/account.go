package logic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/haiyen11231/Internet-download-manager/internal/data_access/database"
)

// tao su tach biet giua layer grpc va logic
// Layer logic nay co the dc dung o layer grpc, kafka hoac cronjob
// khong muon viet logic bi phu thuoc boi logic toi tu dau
type CreateAccountParams struct {
	AccountName string
	Password    string
}

type CreateAccountResponse struct {
	ID        uint64
	AccountName string
}

type Account interface {
	CreateAccount(ctx context.Context, params CreateAccountParams) (CreateAccountResponse, error)
	CreateSession(ctx context.Context, accountName, password string) (CreateAccountResponse, string, error)
}

type account struct {
	goquDatabase             *database.Database
	accountDataAccessor      database.AccountDataAccessor
	accountPasswordDataAccessor database.AccountPasswordDataAccessor
	hashLogic               Hash
}

// constructor
func NewAccount(goquDatabase *database.Database, accountDataAccessor database.AccountDataAccessor, accountPasswordDataAccessor database.AccountPasswordDataAccessor, hashLogic Hash) Account {
	return &account{
		goquDatabase:             goquDatabase,
		accountDataAccessor:      accountDataAccessor,
		accountPasswordDataAccessor: accountPasswordDataAccessor,
		hashLogic:               hashLogic,
	}
}

func (a *account) isAccountUsernameTaken (ctx context.Context, username string) (bool, error) {
	_, err := a.accountDataAccessor.GetAccountByUsername(ctx, username)
	if err != nil {
		// Error when username not exist
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		// Error of connection
		return false, err
	}
	return true, nil
}	

// cam logic nay vao layer grpc handler
func (a *account) CreateAccount(ctx context.Context, params CreateAccountParams) (CreateAccountResponse, error) {
	// flow: create new transaction -> ktra account name taken chua 
	// neu roi thi return alr taken
	// neu chua thi insert account vaof db van dung transaction -> hash password -> insert password vao db van dung transaction

	// ktra transaction co loi k
	var accountID uint64
	txErr := a.goquDatabase.WithTx(func(tx database.Database) error {
		isTaken, err := a.isAccountUsernameTaken(ctx, params.AccountName);
		// tra ve err duy nhat thoi -> new co vde gif xay ra trong transaction thi se rollback lai
		if err != nil {
			return err
		} 
		
		if isTaken {
			return errors.New("account name is already taken")
		}

		// boi tx cung cai dat cac func cua interface db minh da dinh nghia -> co the thay the db cua accountDataAccessor bang transaction db ma k can phai viet lai CreateAccount
		accountID, err = a.accountDataAccessor.WithDatabase(tx).CreateAccount(ctx, &database.Account{
			Username: params.AccountName,
		})
		if err != nil {
			return err
		}

		// khong trace hash tuwf output ra input dc, nhung 2 input giong nhau se co hash giong nhau
		hashedPassword, err := a.hashLogic.Hash(ctx, params.Password)
		if err != nil {
			return err
		}

		if err := a.accountPasswordDataAccessor.WithDatabase(tx).CreateAccountPassword(ctx, &database.AccountPassword{
			UserID: accountID,
			PasswordHash: hashedPassword,
		}); err != nil {
			return err
		}
		return nil
	})

	if txErr != nil {
		return CreateAccountResponse{}, txErr
	}

	return CreateAccountResponse{ID: accountID, AccountName: params.AccountName}, nil
	
}

func (a *account) CreateSession(ctx context.Context, username, password string) (CreateAccountResponse, string, error) {
	// get user by username
	// check password
	// generate token
	return CreateAccountResponse{}, "", nil
}