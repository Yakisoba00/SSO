package auth

import (
	"context"
	"errors"
)

type LoginResult struct {
	AccessToken  string
	RefreshToken string
	Requires2FA  bool
	TwoFAToken   string
}

func (u *UseCase) LoginByEmail(ctx context.Context, email, password string) (*LoginResult, error) {

	return nil, errors.New("invalid email") // Error entity
}

func (u *UseCase) LoginByPhone(ctx context.Context, phone, password string) (*LoginResult, error) {

	return nil, errors.New("invalid phone") // Error entity
}

func (u *UseCase) Logout(ctx context.Context, refreshToken string) error {
	panic("implement me")
}

func (u *UseCase) revokeTokenPair(ctx context.Context, session string) {
	panic("implement me")
}
