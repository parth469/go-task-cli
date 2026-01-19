package task

import (
	"errors"
	"flag"
	"time"

	"github.com/parth469/go-task-cli/internal/resource"
)

func Add(args []string) error {

	addFlag := flag.NewFlagSet("add", flag.ExitOnError)

	var title string
	var description string

	addFlag.StringVar(&title,
		"title",
		"",
		"Create a new task with the given title. E \n Example: -add \"Buy groceries\"",
	)

	addFlag.StringVar(&description,
		"description",
		"",
		"Create a new task with the given description. \n Example: -description \"Buy groceries and milk\"",
	)

	addFlag.Parse(args)

	if title == "" {
		return errors.New("title is required")
	}

	t := time.Now()
	id := t.Unix()

	newTask := Task{
		Id:          int(id),
		Title:       title,
		Description: description,
		CreateAt:    time.Now().UTC(),
		UpdateAt:    time.Now().UTC(),
		Status:      TOOD,
	}

	resource.AppendToFile(newTask)

	return nil

}
