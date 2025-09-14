package database

import (
	"context"
	"log"

	"github.com/doug-martin/goqu/v9"
)

type Account struct {
	UserID   uint64 `sql:"user_id"`
	Username string `sql:"username"`
}

type AccountDataAccessor interface {
	CreateAccount(ctx context.Context, account *Account) (uint64, error)
	GetAccountByID(ctx context.Context, userID uint64) (*Account, error)
	GetAccountByUsername(ctx context.Context, username string) (*Account, error)
	// cai dat transaction - Query password that bai -> cac queries khacs phai revert lai -> cac queris cung thanh cong/cung that bai
	WithDatabase(database Database) AccountDataAccessor
	// transaction db va db bthg se chung 1 so command voi nhau
	// muon dung chung cau lenh GetAccountByUsername hay GetAccountByID trong db goqu bthg hay transaction db thi van muon dung chung cau lenh nhu the
	// -> can interface chung de tru func chung giuwa db bthg va db transaction
}

type accountDataAccessor struct {
	database Database // co th tra ve transaction database -> interact with transaction, thay vi tuong tac voi goqu thi tuong tac voi db bthg
}

func NewAccountDataAccessor(db Database) AccountDataAccessor {
	return &accountDataAccessor{
		database: db,
	}
}

func (a *accountDataAccessor) CreateAccount(ctx context.Context, account *Account) (uint64, error) {
	result, err := a.database.Insert("accounts").Rows(goqu.Record{
		"username":  account.Username,
	}).Executor().ExecContext(ctx)

	if err != nil {
		log.Printf("Failed to create account: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Failed to retrieve last insert ID: %v", err)
		return 0, err
	}
	return uint64(id), nil
}

func (a *accountDataAccessor) GetAccountByID(ctx context.Context, userID uint64) (*Account, error) {
	// implement get account by id in db
	return nil, nil
}		

func (a *accountDataAccessor) GetAccountByUsername(ctx context.Context, username string) (*Account, error) {
	// implement get account by username in db
	return nil, nil
}

// why? khi thay the db cua minh bang transaction db thi van nhan dc doi tuong AccountDataAccessor -> nhung func con lai van dc cai dat dua tren database interface cua minh thay vi dua tren 
// datacbase struct cua goqu
func (a *accountDataAccessor) WithDatabase(database Database) AccountDataAccessor {
	return &accountDataAccessor{
		database: database,
	}
}
