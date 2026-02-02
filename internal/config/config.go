package config

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	HomeDir, _   = os.UserHomeDir()
	TSXVaultPath = filepath.Join(HomeDir, ".vault", "components", "tsx")
	JSXVaultPath = filepath.Join(HomeDir, ".vault", "components", "jsx")
)

type Config struct {
	GitTrackingEnabled bool   `json:"git_tracking_enabled"`
	GitInitialized     bool   `json:"git_initialized"`
	AutoSync           bool   `json:"auto_sync"`
	RemoteUrl          string `json:"remote_url"`
	FirstRunCompleted  bool   `json:"first_run_completed"`
}

func Load() error {
	config, err := os.ReadFile("~/.vault/config.json")
	if err != nil {
		return err
	}
	fmt.Println(config)
	return nil
}
