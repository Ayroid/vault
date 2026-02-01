package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"vault/internal/config"
)

// Error messages
var (
	ErrInvalidComponent = errors.New("invalid component name. component file must have a .tsx or .jsx suffix")
	ErrNameFlagValue    = errors.New("--name requires a value")
)

var vaultPathMap = map[string]string{
	"tsx": config.TSXVaultPath,
	"jsx": config.JSXVaultPath,
}

// NormalizeName converts a component name to lowercase for case-insensitive comparison
func NormalizeName(name string) string {
	return strings.ToLower(name)
}

// IsComponentValid checks if the component has a valid .tsx or .jsx extension
func IsComponentValid(component string) bool {
	component = NormalizeName(component)
	return strings.HasSuffix(component, ".tsx") || strings.HasSuffix(component, ".jsx")
}

// ComponentExists checks if a component exists in the given vault path (case-insensitive)
func ComponentExists(component string, dest string) (string, error) {
	components, err := GetComponents(dest)
	if err != nil {
		return "", fmt.Errorf("error accessing vault: %w", err)
	}

	normalizedComponent := NormalizeName(component)

	for _, comp := range components {
		if NormalizeName(comp) == normalizedComponent {
			return comp, nil
		}
	}

	return "", nil
}

// GetComponentVault returns the vault path for a component based on its extension
func GetComponentVault(component string) (string, error) {
	if !IsComponentValid(component) {
		return "", ErrInvalidComponent
	}

	return vaultPathMap[strings.Split(component, ".")[1]], nil
}

// GetComponents returns a list of component filenames from the given vault path
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

// GetAllComponents returns components from both TSX and JSX vaults
func GetAllComponents() (tsxComponents []string, jsxComponents []string, err error) {
	tsxComponents, err = GetComponents(config.TSXVaultPath)
	if err != nil {
		return nil, nil, err
	}

	jsxComponents, err = GetComponents(config.JSXVaultPath)
	if err != nil {
		return nil, nil, err
	}

	return tsxComponents, jsxComponents, nil
}

// CopyFile copies the contents of src file to dst file
func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

// ParseNameFlag extracts the --name flag value from args, returns remaining args and the name value
func ParseNameFlag(args []string) (remaining []string, name string, err error) {
	remaining = make([]string, 0, len(args))

	for i := 0; i < len(args); i++ {
		if args[i] == "--name" {
			if i+1 >= len(args) {
				return nil, "", ErrNameFlagValue
			}
			name = args[i+1]
			i++ // skip the value
		} else {
			remaining = append(remaining, args[i])
		}
	}

	return remaining, name, nil
}

// PromptYesNo displays a yes/no prompt and returns the user's choice
func PromptYesNo(prompt string, defaultNo bool) bool {
	reader := bufio.NewReader(os.Stdin)
	if defaultNo {
		fmt.Printf("%s (y/N): ", prompt)
	} else {
		fmt.Printf("%s (Y/n): ", prompt)
	}

	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	response = strings.TrimSpace(strings.ToLower(response))
	if response == "" {
		return !defaultNo
	}
	return response == "y" || response == "yes"
}

// PromptText displays a prompt and returns the user's text input
func PromptText(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", prompt)

	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(text)
}

// PromptChoice displays numbered options and returns the user's choice
func PromptChoice(prompt string, options []string) int {
	for i, opt := range options {
		fmt.Printf("  %d. %s\n", i+1, opt)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)

	choice, err := reader.ReadString('\n')
	if err != nil {
		return -1
	}

	choice = strings.TrimSpace(choice)
	for i := range options {
		if choice == fmt.Sprintf("%d", i+1) {
			return i + 1
		}
	}
	return -1
}

// ClearDirectory removes all files (not directories) from the given path
func ClearDirectory(path string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		filePath := path + "/" + entry.Name()
		if err := os.Remove(filePath); err != nil {
			return err
		}
	}

	return nil
}

// RemoveFile removes the specified directory from the given path
func RemoveFile(componentPath string) error {
	if err := os.Remove(componentPath); err != nil {
		return err
	}

	return nil
}
