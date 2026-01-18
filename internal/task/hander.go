package task

import (
	"errors"
	"fmt"
	"os"
)

func Execute() error {
	command := os.Args[1:]

	if len(command) == 0 {
		return errors.New("no command provided. please use the correct syntax. for example: -add \"ask title\"")
	}

	args := command[1:]
	switch os.Args[1] {
	case "add":
		Add(args)
	case "update":
		Update(args)
	default:
		fmt.Println("expected 'add' or 'update' subcommands", os.Args[1])
		os.Exit(1)
	}

	return nil
}
