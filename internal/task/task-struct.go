package task

import "time"

type TaskStatus string

const (
	TOOD      TaskStatus = "Todo"
	INPROCESS TaskStatus = "In-Process"
	DONE      TaskStatus = "Done"
)

type Task struct {
	Id          int
	Title       string
	Description string
	CreateAt    time.Time
	UpdateAt    time.Time
	Status      TaskStatus
}
