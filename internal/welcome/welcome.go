package welcome

import "fmt"

func Show() {
	fmt.Println()
	fmt.Println("=================================")
	fmt.Println("        TASK LIST CLI")
	fmt.Println("=================================")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  task [command]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  add <title>                 Add a new task")
	fmt.Println("  update <id> <title>         Update task title")
	fmt.Println("  delete <id>                 Delete a task")
	fmt.Println()
	fmt.Println("Status Commands:")
	fmt.Println("  mark-todo <id>              Mark task as TODO")
	fmt.Println("  mark-in-progress <id>       Mark task as IN-PROGRESS")
	fmt.Println("  mark-done <id>              Mark task as DONE")
	fmt.Println()
	fmt.Println("List Commands:")
	fmt.Println("  list                        List all tasks")
	fmt.Println("  list-todo                   List TODO tasks")
	fmt.Println("  list-in-progress            List IN-PROGRESS tasks")
	fmt.Println("  list-done                   List DONE tasks")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -h, --help                  Show help information")
	fmt.Println()
}
