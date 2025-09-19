package logic

import (
	"context"
	"errors"

	"github.com/haiyen11231/Internet-download-manager/internal/configs"
	"golang.org/x/crypto/bcrypt"
)

type Hash interface {
	Hash(ctx context.Context, data string) (string, error)
	IsHashEqual(ctx context.Context, data, hashed string) (bool, error)
}

type hash struct {
	accountConfig configs.Account
}

func NewHash(accountConfig configs.Account) Hash {
	return &hash{
		accountConfig: accountConfig,
	}
}

func (h hash) Hash(ctx context.Context, data string) (string, error) {
	// implement hash function
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data), h.accountConfig.HashCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (h hash) IsHashEqual(ctx context.Context, data, hashed string) (bool, error) {
	// implement check hash function
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(data)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}