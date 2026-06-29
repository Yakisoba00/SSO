package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		App    App
		Log    Log
		Tokens Tokens
		Otp    Otp
	}

	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}

	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	Tokens struct {
		AccessExpiry  time.Duration `env:"ACCESS_TOKEN_EXPIRY" envDefault:"15m"`
		RefreshExpiry time.Duration `env:"REFRESH_TOKEN_EXPIRY" envDefault:"24h"`
		Issuer        string        `env:"TOKEN_ISSUER,required"`
		SecretSeed    string        `env:"SECRET_SEED,required"`
	}

	Otp struct {
		CodeLength  int           `env:"OTP_CODE_LENGTH" envDefault:"6"`
		MaxAttempts int           `env:"OTP_MAX_ATTEMPTS" envDefault:"3"`
		CodeTTL     time.Duration `env:"OTP_CODE_TTL" envDefault:"5m"`
		Block       time.Duration `env:"OTP_BLOCK" envDefault:"10m"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %v", err)
	}
	return cfg, nil
}
