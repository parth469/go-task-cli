package task

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("fail to parse id")
		}
		Update(args[1:], id)
	case "delete":
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("fail to parse id")
		}
		Delete(id)
	case "list":
		List(args)

	default:
		fmt.Println("expected 'add' or 'update' subcommands", os.Args[1])
		os.Exit(1)
	}

	return nil
}
