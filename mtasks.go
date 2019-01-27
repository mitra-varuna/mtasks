package mtasks

// Tasks represent a common task that needs to be done
type Task struct {
	Title       string
	Description string
	IsDone      bool
}

// Tasker represents an interface which manages task
type Tasker interface {
	Create() error
	Edit() (bool, error)
	MarkDone()
}
