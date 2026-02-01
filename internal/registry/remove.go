package registry

import (
	"errors"
	"fmt"
	"path/filepath"
	"slices"

	"vault/internal/utils"
)

func HandleDelete(args []string) error {
	if len(args) < 1 {
		return errors.New("usage: vault delete <component>")
	}

	components, err := ListFiles()
	if err != nil {
		return err
	}

	if !slices.Contains(components, args[0]) {
		return fmt.Errorf("component does not exist in vault: %s", args[0])
	}

	return deleteComponent(args[0])
}

func deleteComponent(component string) error {
	fmt.Printf("This will delete %s component from the vault.\n", component)

	if !utils.PromptYesNo("Are you sure you want delete?", true) {
		fmt.Println("Operation aborted")
		return nil
	}

	componentVault, err := utils.GetComponentVault(component)
	if err != nil {
		return err
	}

	componentPath := filepath.Join(componentVault, component)

	if err := utils.RemoveFile(componentPath); err != nil {
		return fmt.Errorf("failed to delete component: %w", err)
	}

	fmt.Println("Component deleted successfully")
	return nil
}
