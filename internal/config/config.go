package config

import (
	"os"
	"path/filepath"
)

var (
	HomeDir, _   = os.UserHomeDir()
	TSXVaultPath = filepath.Join(HomeDir, ".vault", "components", "tsx")
	JSXVaultPath = filepath.Join(HomeDir, ".vault", "components", "jsx")
)
