package registry

import (
	"errors"
	"fmt"
	"os"
	"vault/internal/config"
)

func HandleList(args []string) error {
	if len(args) != 0 {
		return errors.New("list command takes no arguments")
	}

	files, err := ListFiles()
	if err != nil {
		return errors.New("error listing components")
	}

	for _, file := range files {
		fmt.Println(file)
	}
	return nil
}

func ListFiles() ([]string, error) {

	entries, err := os.ReadDir(config.VaultPath)

	if err != nil {
		return nil, err
	}

	files := make([]string, 0, len(entries))

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		files = append(files, entry.Name())
	}

	return files, nil
}
