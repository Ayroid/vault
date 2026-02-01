package registry

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"vault/internal/utils"
)

func HandleSave(args []string) error {
	if len(args) < 1 {
		return errors.New("usage: vault save <file> [--name <newname>]")
	}

	remaining, name, err := utils.ParseNameFlag(args[1:])
	if err != nil {
		return err
	}

	if len(remaining) > 0 {
		return fmt.Errorf("unknown flag: %s", remaining[0])
	}

	return saveComponent(args[0], name)
}

func saveComponent(src string, customName string) error {
	absSrc, err := filepath.Abs(src)
	if err != nil {
		return err
	}

	if _, err := os.Stat(absSrc); err != nil {
		return err
	}

	component := filepath.Base(absSrc)
	if customName != "" {
		component = customName
	}

	if !utils.IsComponentValid(component) {
		return utils.ErrInvalidComponent
	}

	vaultPath, err := utils.GetComponentVault(component)
	if err != nil {
		return err
	}

	existingComponent, err := utils.ComponentExists(component, vaultPath)
	if err != nil {
		return err
	}

	if existingComponent != "" {
		return fmt.Errorf("component already exists (case-insensitive match: %s)", existingComponent)
	}

	destinationPath := filepath.Join(vaultPath, component)
	return utils.CopyFile(absSrc, destinationPath)
}
