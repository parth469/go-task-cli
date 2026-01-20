package task

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"

	"github.com/parth469/go-task-cli/internal/resource"
)

func Update(args []string, id int) error {
	updateFlag := flag.NewFlagSet("update", flag.ExitOnError)

	var title string
	var description string

	updateFlag.StringVar(
		&title,
		"title",
		"",
		"Update the title of the task. \nExample usage: -title \"New title for the task\". If omitted, the title will remain unchanged.",
	)

	updateFlag.StringVar(
		&description,
		"description",
		"",
		"Update the description of the task. \nExample usage: -description \"Updated task details here\". If omitted, the description will remain unchanged.",
	)

	updateFlag.Parse(args)

	updateFlag.Parse(args)

	if title == "" && description == "" {
		return errors.New("nothing to update: provide -title and/or -description")
	}

	jsonData, err := resource.ListAll()
	if err != nil {
		return fmt.Errorf("read tasks: %w", err)
	}

	var taskLists []Task
	if err := json.Unmarshal(jsonData, &taskLists); err != nil {
		return fmt.Errorf("decode tasks json: %w", err)
	}

	recordFound := false
	for i, v := range taskLists {
		if v.Id == id {
			recordFound = true
			if title != "" {
				taskLists[i].Title = title
			}
			if description != "" {
				taskLists[i].Description = description
			}
			break
		}
	}

	if !recordFound {
		return fmt.Errorf("task not found: id=%d", id)
	}

	byteData, err := json.Marshal(taskLists)
	if err != nil {
		return fmt.Errorf("encode tasks json: %w", err)
	}

	if err := resource.WriteToFile(byteData); err != nil {
		return fmt.Errorf("write tasks file: %w", err)
	}

	return nil
}
