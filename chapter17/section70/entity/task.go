package entity

import "time"

// 자체 타입을 지정하여, 잘못된 대입 방지
type TaskID int64
type TaskStatus string

const (
	TaskStatusTodo TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone TaskStatus = "done"
)

type Task struct {
	ID		TaskID		`json:"id"`
	Title	string		`json:"title"`
	Status	TaskStatus	`json:"status"`
	Created	time.Time	`json:"created"`
}

type Tasks []*Task