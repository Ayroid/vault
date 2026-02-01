package registry

import (
	"errors"
	"fmt"

	"vault/internal/config"
	"vault/internal/utils"
)

func HandleClear(args []string) error {
	if len(args) != 0 {
		return errors.New("clear command takes no arguments")
	}

	components, err := ListFiles()
	if err != nil {
		return err
	}

	if len(components) == 0 {
		fmt.Println("vault is already empty")
		return nil
	}

	fmt.Printf("This will delete %d component(s) from the vault.\n", len(components))

	if !utils.PromptYesNo("Are you sure you want to clear the vault?", true) {
		fmt.Println("Operation aborted")
		return nil
	}

	if err := utils.ClearDirectory(config.TSXVaultPath); err != nil {
		return fmt.Errorf("failed to clear TSX vault: %w", err)
	}

	if err := utils.ClearDirectory(config.JSXVaultPath); err != nil {
		return fmt.Errorf("failed to clear JSX vault: %w", err)
	}

	fmt.Println("Vault cleared successfully")
	return nil
}
