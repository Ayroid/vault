package registry

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"gocopy/internal/config"
)

func HandleSave(args []string) error {
	if len(args) < 1 {
		return errors.New("usage gocopy save <filename>")
	}

	src := args[0]

	return CopyFile(src)
}

func CopyFile(src string) error {
	sourceFile, err := os.Open(src)

	if err != nil {
		return err
	}
	defer sourceFile.Close()

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
