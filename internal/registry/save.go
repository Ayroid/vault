package registry

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"

	"vault/internal/config"
)

type SaveOptions struct {
	Name string
}

func parseSaveArgs(args []string) (string, SaveOptions, error) {
	if len(args) < 1 {
		return "", SaveOptions{}, errors.New("usage: vault save <file> [--name <newname>]")
	}

	src := args[0]
	opts := SaveOptions{}

	i := 1
	for i < len(args) {
		switch args[i] {
		case "--name":
			if i+1 >= len(args) {
				return "", SaveOptions{}, errors.New("--name requires a value")
			}
			opts.Name = args[i+1]
			i += 2
		default:
			return "", SaveOptions{}, fmt.Errorf("unknown flag: %s", args[i])
		}
	}
	return src, opts, nil
}

func HandleSave(args []string) error {
	src, opts, err := parseSaveArgs(args)
	if err != nil {
		return err
	}
	return CopyFile(src, opts)
}

func filenameExists(filename string) (bool, error) {
	files, err := ListFiles()
	if err != nil {
		return false, errors.New("error saving component")
	}

	filenameExist := slices.Contains(files, filename)

	if filenameExist {
		return true, nil
	}

	return false, nil
}

func CopyFile(src string, opts SaveOptions) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	filename := filepath.Base(sourceFile.Name())

	if opts.Name != "" {
		filename = opts.Name
	}

	exists, err := filenameExists(filename)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("component with same name already exists in vault")
	}

	dest := filepath.Join(config.VaultPath, filename)

	destinationFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}
