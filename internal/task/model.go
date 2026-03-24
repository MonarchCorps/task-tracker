package task

type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in-progress"
	Done       Status = "done"
)

type Task struct {
	ID          int
	Description string
	Status      Status
}

type TaskFilter struct {
	Statuses []Status
}
