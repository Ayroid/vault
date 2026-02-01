package main

import (
	"fmt"
	"log"
	"os"

	"vault/internal/registry"
)

type Command string
type CommandHandler func(args []string) error

const (
	CmdSave   Command = "save"
	CmdList   Command = "list"
	CmdGet    Command = "get"
	CmdClear  Command = "clear"
	CmdDelete Command = "delete"
	CmdHelp   Command = "help"
)

var CommandHandlers = map[Command]CommandHandler{
	CmdSave:   registry.HandleSave,
	CmdList:   registry.HandleList,
	CmdGet:    registry.HandleGet,
	CmdClear:  registry.HandleClear,
	CmdDelete: registry.HandleDelete,
	CmdHelp:   handleHelp,
}

func handleHelp(args []string) error {
	help := `Vault - A lightweight React/Next.js component registry CLI

Usage:
  vault <command> [arguments]

Commands:
  save <component> [--name <newname>] Save a component to the vault
  get <component> [--name <name>]     Retrieve a component from the vault
  list                                List all components in the vault
  clear                               Remove all components from the vaultdelete
  delete <component>                  Delete specified component from the vault
  help, --help, -h                    Show this help message

Examples:
  vault save Button.tsx               Save Button.tsx to the vault
  vault save Button.tsx --name Btn.tsx  Save with a different name
  vault get Button.tsx                Copy Button.tsx to your project
  vault get Button.tsx --name MyButton.tsx  Copy with a different name
  vault list                          Show all saved components
  vault delete Button.tsx             Delete Button.tsx from the vault
  vault clear                         Clear the vault (with confirmation)

Storage:
  Components are stored in ~/.vault/components/ organized by type (tsx/jsx).`

	fmt.Println(help)
	return nil
}

func isHelpFlag(arg string) bool {
	return arg == "--help" || arg == "-h" || arg == "help"
}

func getCommandHandler(cmd Command) (CommandHandler, error) {
	handler, ok := CommandHandlers[cmd]
	if !ok {
		return nil, fmt.Errorf("unknown command: %s\nRun 'vault --help' for usage", cmd)
	}
	return handler, nil
}

func main() {
	if err := registry.SetupVault(); err != nil {
		log.Fatal(err)
	}

	args := os.Args[1:]

	if len(args) == 0 || isHelpFlag(args[0]) {
		handleHelp(nil)
		return
	}

	cmd := Command(args[0])
	cmdArgs := args[1:]

	handler, err := getCommandHandler(cmd)
	if err != nil {
		log.Fatal(err)
	}

	if err := handler(cmdArgs); err != nil {
		log.Fatal(err)
	}
}
