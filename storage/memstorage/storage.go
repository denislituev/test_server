package memstorage

import (
	"strconv"
	"test_server/domain"
)

type MemStorage struct{
	tasks domain.Tasks
	id int64
}

func NewMemStorage() *MemStorage{
	return &MemStorage{}
}

func (ms *MemStorage) GetAllTasks() (domain.Tasks, error){
	return ms.tasks, nil
}

func (ms *MemStorage) CreateTask(task domain.Task) (domain.Task, error){
	ms.id++
	task.Id = strconv.FormatInt(ms.id, 10)
	ms.tasks.TaskList = append(ms.tasks.TaskList, task)
	return task, nil
}