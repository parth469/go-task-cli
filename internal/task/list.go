package task

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/parth469/go-task-cli/internal/resource"
)

func PrintJSON(data []Task) {
	if len(data) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	formattedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}
	fmt.Println(string(formattedJSON))
}

func List(args []string) error {
	list := flag.NewFlagSet("list", flag.ExitOnError)

	var rawStatus string
	var data []Task
	list.StringVar(&rawStatus, "status", "", "Please provide status one of in-progress, done , todo")

	list.Parse(args)

	jsonData, err := resource.ListAll()

	if err != nil {
		return fmt.Errorf("fail to get list data %w", err)
	}

	if err := json.Unmarshal(jsonData, &data); err != nil {
		return fmt.Errorf("parse json: %w", err)
	}

	if rawStatus == "" {
		PrintJSON(data)
		return nil
	}

	filterData := []Task{}

	var status TaskStatus
	switch rawStatus {
	case "todo":
		status = TOOD
	case "in-progress":
		status = INPROCESS
	case "done":
		status = DONE
	}

	for _, v := range data {
		if v.Status == status {
			filterData = append(filterData, v)
		}
	}

	PrintJSON(filterData)
	return nil
}
