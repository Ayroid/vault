package registry

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"vault/internal/utils"
)

func HandleGet(args []string) error {
	if len(args) < 1 {
		return errors.New("usage: vault get <file> [--name <newname>]")
	}

	remaining, customName, err := utils.ParseNameFlag(args[1:])
	if err != nil {
		return err
	}

	if len(remaining) > 0 {
		return fmt.Errorf("unknown flag: %s", remaining[0])
	}

	return getComponent(args[0], customName)
}

func findComponentsDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	possiblePaths := []string{
		filepath.Join(cwd, "components"),
		filepath.Join(cwd, "app", "components"),
		filepath.Join(cwd, "src", "components"),
	}

	for _, path := range possiblePaths {
		if info, err := os.Stat(path); err == nil && info.IsDir() {
			return path, nil
		}
	}

	return "", nil
}

func getComponent(component string, customName string) error {
	if !utils.IsComponentValid(component) {
		return errors.New("invalid component name or component does not exist in vault")
	}

	vaultPath, err := utils.GetComponentVault(component)
	if err != nil {
		return err
	}

	existingComponent, err := utils.ComponentExists(component, vaultPath)
	if err != nil {
		return err
	}

	if existingComponent == "" {
		return fmt.Errorf("component does not exist in the vault: %s", component)
	}

	sourcePath := filepath.Join(vaultPath, existingComponent)

	targetName := existingComponent
	if customName != "" {
		targetName = customName
		if !utils.IsComponentValid(targetName) {
			return utils.ErrInvalidComponent
		}
	}

	componentsDir, err := findComponentsDir()
	if err != nil {
		return err
	}

	if componentsDir == "" {
		if !utils.PromptYesNo("No components directory found. Create new components/vault directory?", true) {
			fmt.Println("Operation aborted")
			return nil
		}

		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		componentsDir = filepath.Join(cwd, "components", "vault")
		if err := os.MkdirAll(componentsDir, 0755); err != nil {
			return fmt.Errorf("failed to create components directory: %w", err)
		}
		fmt.Printf("Created directory: %s\n", componentsDir)
	}

	destPath := filepath.Join(componentsDir, targetName)

	if _, err := os.Stat(destPath); err == nil {
		fmt.Printf("File '%s' already exists in %s\n", targetName, componentsDir)
		fmt.Println("Choose an option:")

		choice := utils.PromptChoice("Enter choice (1/2/3): ", []string{
			"Replace existing file",
			"Rename and save with a new name",
			"Abort",
		})

		switch choice {
		case 1:
			// Replace - continue with copy
		case 2:
			newName := utils.PromptText("Enter new name (with .tsx or .jsx extension)")
			if newName == "" {
				return errors.New("invalid name provided")
			}
			if !utils.IsComponentValid(newName) {
				return utils.ErrInvalidComponent
			}
			destPath = filepath.Join(componentsDir, newName)
			targetName = newName
		default:
			fmt.Println("Operation aborted")
			return nil
		}
	}

	if err := utils.CopyFile(sourcePath, destPath); err != nil {
		return fmt.Errorf("failed to copy component: %w", err)
	}

	fmt.Printf("Component '%s' copied to %s\n", targetName, destPath)
	return nil
}
