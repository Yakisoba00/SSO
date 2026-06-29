package auth

import "context"

func (u *UseCase) RefreshTokenPair(ctx context.Context, refreshToken string) (access, refresh string, err error) {
	panic("implement me")
}
