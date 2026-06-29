package pasetov4

import "errors"

var (
	ErrInvalidToken  = errors.New("invalid token")
	ErrInvalidKey    = errors.New("invalid public or private key")
	ErrClaimSetting  = errors.New("failed to set token claim")
	ErrTokenCreation = errors.New("failed to create token")
)
