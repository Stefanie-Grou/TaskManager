package models

type task struct {
	TaskName        string
	TaskDescription string
	Done            bool
}

func CreateTask(TaskName string, TaskDescription string, Done bool) *task {
	return &task{
		TaskName:        TaskName,
		TaskDescription: TaskDescription,
		Done:            Done,
	}
}
