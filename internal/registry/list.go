package registry

import (
	"errors"
	"fmt"
	"os"

	"vault/internal/config"
)

func treeDisplay(components []string) {
	for index, file := range components {
		prefix := "├── "
		if index == len(components)-1 {
			prefix = "└── "
		}
		fmt.Println(prefix + file)
	}
}

func GetComponents(vaultPath string) ([]string, error) {
	componentFiles, err := os.ReadDir(vaultPath)
	if err != nil {
		return nil, err
	}
	components := make([]string, 0, len(componentFiles))

	for _, component := range componentFiles {
		if component.IsDir() {
			continue
		}
		components = append(components, component.Name())
	}
	return components, nil
}

func HandleList(args []string) error {
	if len(args) != 0 {
		return errors.New("list command takes no arguments")
	}

	tsxComponents, err := GetComponents(config.TSXVaultPath)
	if err != nil {
		return errors.New("error listing components")
	}

	jsxComponents, err := GetComponents(config.JSXVaultPath)
	if err != nil {
		return errors.New("error listing components")
	}

	if len(jsxComponents)+len(tsxComponents) == 0 {
		fmt.Println("vault is empty")
		return nil
	}

	tsxComponentsLength := len(tsxComponents)
	if tsxComponentsLength > 0 {
		fmt.Println("TSX Components")
		treeDisplay(tsxComponents)
	}

	jsxComponentsLength := len(jsxComponents)
	if jsxComponentsLength > 0 {
		fmt.Println("JSX Components")
		treeDisplay(jsxComponents)
	}

	return nil
}

func ListFiles() ([]string, error) {

	tsxComponents, err := GetComponents(config.TSXVaultPath)
	if err != nil {
		return nil, err
	}

	jsxComponents, err := GetComponents(config.JSXVaultPath)
	if err != nil {
		return nil, err
	}

	components := append(tsxComponents, jsxComponents...)

	return components, nil
}
