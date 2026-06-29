package auth

import (
	"context"
	"sso/pkg/pasetov4"
)

type UseCase struct {
	tokenProvider *pasetov4.Provider
}

func New(tokenProvider *pasetov4.Provider) *UseCase {
	return &UseCase{
		tokenProvider: tokenProvider,
	}
}

func (u *UseCase) issueTokenPair(ctx context.Context) (access, refresh string, err error) {
	accessClaims := map[string]interface{}{"user_id": 1, "role": "user.Role"}
	refreshClaims := map[string]interface{}{"user_id": 1, "session_id": "uuid.NewString()"}

	return u.tokenProvider.GeneratePair(accessClaims, refreshClaims)
}
