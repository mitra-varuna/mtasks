package mtasks

// Tasks represent a common task that needs to be done
type Task struct {
	Title       string
	Description string
}

// TaskManager represents an interface which manages task
type TaskManager interface {
	Create() (Task, error)
	Edit(t Task) (bool, error)
	MarkDone(t Task)
}
