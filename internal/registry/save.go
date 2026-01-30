package registry

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"

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

func validComponent(component string) bool {
	component = strings.ToLower(component)
	return strings.HasSuffix(component, ".tsx") || strings.HasSuffix(component, ".jsx")
}

func componentExists(component string) (bool, error) {
	tsxComponent := strings.HasSuffix(component, ".tsx")

	var vaultPath string

	if tsxComponent {
		vaultPath = config.TSXVaultPath
	} else {
		vaultPath = config.JSXVaultPath
	}

	files, err := GetComponents(vaultPath)
	if err != nil {
		return false, errors.New("error saving component")
	}

	filenameExist := slices.Contains(files, component)

	if filenameExist {
		return true, nil
	}

	return false, nil
}

func CopyFile(src string, opts SaveOptions) error {
	src = strings.ToLower(src)

	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	component := filepath.Base(sourceFile.Name())

	if opts.Name != "" {
		component = opts.Name
	}

	if !validComponent(component) {
		return errors.New("invalid file. must be a component file with .tsx or .jsx extension")
	}

	exists, err := componentExists(component)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("component with same name already exists in vault")
	}

	tsxComponent := strings.HasSuffix(component, ".tsx")

	var dest string

	if tsxComponent {
		dest = filepath.Join(config.TSXVaultPath, component)
	} else {
		dest = filepath.Join(config.JSXVaultPath, component)
	}

	destinationFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}
