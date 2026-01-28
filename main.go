package main

import (
	"fmt"
	"log"
	"os"

	"gocopy/internal/registry"
)

type Command string
type CommandHandler func(args []string) error

const (
	CmdSave Command = "save"
	CmdList Command = "list"
)

var CommandHandlers = map[Command]CommandHandler{
	CmdSave: registry.HandleSave,
	CmdList: registry.HandleList,
}

func getCommandHandler(cmd Command) (CommandHandler, error) {
	handler, ok := CommandHandlers[cmd]
	if !ok {
		return nil, fmt.Errorf("unknown command %s", cmd)
	}
	return handler, nil
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("no command provided")
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
