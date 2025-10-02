package database

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/haiyen11231/Internet-download-manager/internal/utils"
	"go.uber.org/zap"
)

var (
	TabNameTokenPublicKeys          = goqu.T("token_public_keys")
)

const (
	ColNameTokenPublicKeysID        = "key_id"
	ColNameTokenPublicKeysPublicKey = "public_key"
)

type TokenPublicKey struct {
	KeyID     uint64 `db:"key_id" goqu:"skipinsert,skipupdate"`
	PublicKey string `db:"public_key"`
}

type TokenPublicKeyDataAccessor interface {
	CreatePublicKey(ctx context.Context, tokenPublicKey TokenPublicKey) (uint64, error)
	GetPublicKey(ctx context.Context, id uint64) (TokenPublicKey, error)
	WithDatabase(database Database) TokenPublicKeyDataAccessor
}

type tokenPublicKeyDataAccessor struct {
	database Database
	logger   *zap.Logger
}

func NewTokenPublicKeyDataAccessor(database *goqu.Database, logger *zap.Logger) TokenPublicKeyDataAccessor {
	return &tokenPublicKeyDataAccessor{
		database: database,
		logger:   logger,
	}
}	

func (t tokenPublicKeyDataAccessor) CreatePublicKey(ctx context.Context, tokenPublicKey TokenPublicKey) (uint64, error) {
	logger := utils.LoggerWithContext(ctx, t.logger)
	result, err := t.database.Insert(TabNameTokenPublicKeys).Rows(goqu.Record{
		ColNameTokenPublicKeysPublicKey: tokenPublicKey.PublicKey,
	}).Executor().ExecContext(ctx)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to create token public key")
		return 0, status.Errorf(codes.Internal, "failed to create token public key: %+v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get last insert id")
		return 0, status.Errorf(codes.Internal, "failed to get last insert id: %+v", err)
	}
	return uint64(id), nil
}

func (t tokenPublicKeyDataAccessor) GetPublicKey(ctx context.Context, id uint64) (TokenPublicKey, error) {
	logger := utils.LoggerWithContext(ctx, t.logger).With(zap.Uint64("id", id))

	tokenPublicKey := TokenPublicKey{}
	found, err := t.database.Select().From(TabNameTokenPublicKeys).Where(goqu.Ex{
		ColNameTokenPublicKeysID: id,
	}).Executor().ScanStructContext(ctx, &tokenPublicKey)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get token public key")
		return TokenPublicKey{}, status.Errorf(codes.Internal, "failed to get token public key: %+v", err)
	}

	if !found {
		logger.Warn("token public key not found")
		return TokenPublicKey{}, sql.ErrNoRows
	}
	
	return tokenPublicKey, nil
}

func (t tokenPublicKeyDataAccessor) WithDatabase(database Database) TokenPublicKeyDataAccessor {
	t.database = database
	return t
}