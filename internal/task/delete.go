package task

import (
	"encoding/json"
	"fmt"

	"github.com/parth469/go-task-cli/internal/resource"
)

func Delete(id int) error {
	var tasks []Task
	jsonData, err := resource.ListAll()

	if err != nil {
		return fmt.Errorf("fail to parse data %w", err)
	}

	json.Unmarshal(jsonData, &tasks)

	updateData := []Task{}

	findRecord := false
	for _, v := range tasks {
		if v.Id == id {
			findRecord = true
			continue
		}
		updateData = append(updateData, v)
	}

	if !findRecord {
		return fmt.Errorf("record not found")
	}

	updateByte, err := json.Marshal(updateData)

	if err != nil {
		return fmt.Errorf("fail to marsh json %w", err)
	}

	resource.WriteToFile(updateByte)

	return nil
}
