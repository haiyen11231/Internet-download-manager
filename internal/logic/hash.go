package logic

import (
	"context"

	"github.com/haiyen11231/Internet-download-manager/internal/configs"
	"golang.org/x/crypto/bcrypt"
)

type Hash interface {
	Hash(ctx context.Context, data string) (string, error)
	IsHashEqual(ctx context.Context, data, hash string) (bool, error)
}

type hash struct {
	accountConfig *configs.AccountConfig
}

func NewHash(accountConfig *configs.AccountConfig) Hash {
	return &hash{
		accountConfig: accountConfig,
	}
}

func (h *hash) Hash(ctx context.Context, data string) (string, error) {
	// implement hash function
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data), h.accountConfig.HashCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (h *hash) IsHashEqual(ctx context.Context, data, hash string) (bool, error) {
	// implement check hash function
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}