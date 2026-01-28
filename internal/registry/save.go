package registry

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"slices"

	"vault/internal/config"
)

func HandleSave(args []string) error {
	if len(args) < 1 {
		return errors.New("usage vault save <filename>")
	}

	src := args[0]

	return CopyFile(src)
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

func CopyFile(src string) error {
	sourceFile, err := os.Open(src)

	if err != nil {
		return err
	}
	defer sourceFile.Close()

	fileExists, err := filenameExists(sourceFile.Name())
	if err != nil {
		return err
	}

	if fileExists {
		return errors.New("component with same name already exists in vault")
	}

	dest := config.VaultPath + sourceFile.Name()

	err = os.MkdirAll(filepath.Dir(dest), 0755)
	if err != nil {
		return err
	}

	destinationFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}
