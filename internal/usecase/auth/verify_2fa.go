package auth

import (
	"context"
	"errors"
)

func (u *UseCase) Verify2FA(ctx context.Context, twoFAToken, code string) (access, refresh string, err error) {

	return "", "", errors.New("invalid 2FA")
}
