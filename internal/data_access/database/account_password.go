package database

import (
	"context"

	"github.com/doug-martin/goqu/v9"
)

type AccountPassword struct {
	UserID       uint64 `sql:"of_user_id"`
	PasswordHash string `sql:"password_hash"`
}

type AccountPasswordDataAccessor interface {
	CreateAccountPassword(ctx context.Context, accountPassword *AccountPassword) error
	UpdateAccountPassword(ctx context.Context, accountPassword *AccountPassword) error
	WithDatabase(database Database) AccountPasswordDataAccessor
}

type accountPasswordDataAccessor struct {
	// database *goqu.Database
	database Database
}

func NewAccountPasswordDataAccessor(db *goqu.Database) AccountPasswordDataAccessor {
	return &accountPasswordDataAccessor{
		database: db,
	}
}

func (a *accountPasswordDataAccessor) CreateAccountPassword(ctx context.Context, accountPassword *AccountPassword) error {
	_, err := a.database.Insert("account_passwords").Rows(accountPassword).Executor().ExecContext(ctx)
	return err
}

func (a *accountPasswordDataAccessor) UpdateAccountPassword(ctx context.Context, accountPassword *AccountPassword) error {}

func (a *accountPasswordDataAccessor) WithDatabase(database Database) AccountPasswordDataAccessor {
	return &accountPasswordDataAccessor{
		database: database.(*goqu.Database),
	}
}