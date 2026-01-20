package task

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/parth469/go-task-cli/internal/resource"
)

func StatusUpdate(id int, rawStatus string) error {
	var status TaskStatus
	switch rawStatus {
	case "done":
		status = DONE
	case "todo":
		status = TOOD
	case "in-progress":
		status = INPROCESS
	default:
		return errors.New("please provide valid status")
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
			taskLists[i].Status = status
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
