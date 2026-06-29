package app

import (
	"fmt"
	"sso/config"
	"sso/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	l.Info("Running apps...")
	l.Info(fmt.Sprintf("Name: %s, Version: %s", cfg.App.Name, cfg.App.Version))
}
