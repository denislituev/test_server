package domain

type Tasks struct {
	TaskList []Task
}

type Task struct {
	Id          string
	Name        string
	Description string
}
