package otp

import (
	"sso/config"
	"time"
)

type UseCase struct {
	codeLength  int
	maxAttempts int
	codeTTL     time.Duration
	block       time.Duration

	// storage
}

func New(
	cfg config.Otp,
) (*UseCase, error) {

	return &UseCase{
		codeLength:  cfg.CodeLength,
		maxAttempts: cfg.MaxAttempts,
		codeTTL:     cfg.CodeTTL,
		block:       cfg.Block,
	}, nil
}
