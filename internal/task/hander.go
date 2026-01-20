package task

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Execute() error {
	if len(os.Args) < 2 {
		return errors.New("no command provided")
	}

	action := os.Args[1]
	args := os.Args[2:]

	switch action {
	case "add":
		return wrap("create record", Add(args))

	case "update":
		id, rest, err := parseID(args)
		if err != nil {
			return err
		}
		return wrap("update record", Update(rest, id))

	case "delete":
		id, _, err := parseID(args)
		if err != nil {
			return err
		}
		return wrap("delete record", Delete(id))

	case "mark":
		id, _, err := parseID(args)
		status := os.Args[3]

		if status == "" {
			return errors.New("please provide status")
		}
		if err != nil {
			return err
		}
		return wrap("mark record", StatusUpdate(id, status))

	case "list":
		return List(args)

	default:
		return fmt.Errorf("unknown command: %s", action)
	}
}

func parseID(args []string) (int, []string, error) {
	if len(args) == 0 {
		return 0, nil, errors.New("missing id")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, nil, errors.New("invalid id")
	}

	return id, args[1:], nil
}

func wrap(action string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("failed to %s: %w", action, err)
}
