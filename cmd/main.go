package main

import (
	"log"
	"os"

	"github.com/parth469/go-task-cli/internal/resource"
	"github.com/parth469/go-task-cli/internal/task"
	"github.com/parth469/go-task-cli/internal/welcome"
)

func main() {
	welcome.Show()

	if err := resource.FileCheck(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if err := task.Execute(); err != nil {
		log.Fatal(err)
	}
}
