package registry

import (
	"errors"
	"fmt"

	"vault/internal/utils"
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

func HandleList(args []string) error {
	if len(args) != 0 {
		return errors.New("list command takes no arguments")
	}

	tsxComponents, jsxComponents, err := utils.GetAllComponents()
	if err != nil {
		return errors.New("error listing components")
	}

	if len(jsxComponents)+len(tsxComponents) == 0 {
		fmt.Println("vault is empty")
		return nil
	}

	if len(tsxComponents) > 0 {
		fmt.Println("TSX Components")
		treeDisplay(tsxComponents)
	}

	if len(jsxComponents) > 0 {
		fmt.Println("JSX Components")
		treeDisplay(jsxComponents)
	}

	return nil
}

func ListFiles() ([]string, error) {
	tsxComponents, jsxComponents, err := utils.GetAllComponents()
	if err != nil {
		return nil, err
	}

	return append(tsxComponents, jsxComponents...), nil
}
