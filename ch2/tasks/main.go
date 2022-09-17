package main

import "time"

type TaskID int64
type TaskStatus string

const (
	TaskStatusTodo  TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone  TaskStatus = "done"
)

type Task struct {
	ID      TaskID     `json:"id"`
	Title   string     `json:"title"`
	Status  TaskStatus `json:"status"`
	Created time.Time  `json:"created"`
}

type Tasks []*Task

func main() {
	var id int = 1
	// _ = Task{ID: id} // cannot use id (type int) as type TaskID in field value
	_ = Task{ID: TaskID(id)}

	_ = Task{ID: 1}
}
