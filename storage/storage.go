package storage

import "test_server/domain"

type Storage interface {
	GetAllTasks() (domain.Tasks, error)
	CreateTask(task domain.Task) (domain.Task, error)
}