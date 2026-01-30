package registry

import (
	"os"

	"vault/internal/config"
)

func SetupVault() error {
	vaults := []string{
		config.TSXVaultPath,
		config.JSXVaultPath,
	}

	for _, vault := range vaults {
		if err := os.MkdirAll(vault, 0755); err != nil {
			return err
		}
	}
	return nil
}
