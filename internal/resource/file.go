package resource

import (
	"encoding/json"
	"fmt"
	"os"
)

var filePath = "db.json"

func FileCheck() error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDONLY, 0644)

	if err != nil {
		return fmt.Errorf("file to create/open file %w", err)
	}

	defer file.Close()

	return nil
}

func AppendToFile(v any) error {
	var list []any

	data, _ := os.ReadFile(filePath)
	if len(data) > 0 {
		if err := json.Unmarshal(data, &list); err != nil {
			return err
		}
	}

	// Append
	list = append(list, v)

	// Write back
	out, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, out, 0644)
}
