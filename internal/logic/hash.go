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
	authConfig configs.Auth
}

func NewHash(authConfig configs.Auth) Hash {
	return &hash{
		authConfig: authConfig,
	}
}

func (h hash) Hash(_ context.Context, data string) (string, error) {
	// implement hash function
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data), h.authConfig.Hash.HashCost)
	if err != nil {
		return "", status.Errorf(codes.Internal, "failed to hash data: %+v", err)
	}
	return string(hashedPassword), nil
}

func (h hash) IsHashEqual(_ context.Context, data, hashed string) (bool, error) {
	// implement check hash function
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(data)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, status.Errorf(codes.Internal, "failed to check if data equal hash: %+v", err)
	}
	return true, nil
}