package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test_server/domain"
	"test_server/storage"
)

type TaskHandlers struct{
	storage storage.Storage
}

func NewTaskHandlers(store storage.Storage) TaskHandlers{
	return TaskHandlers{
		storage: store,
	}
}

func (th TaskHandlers) GetAllTasks(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tasks, err := th.storage.GetAllTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Cannot get tasks due to: ", err.Error())
	}


	data, err := json.Marshal(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Cannot marshal tasks due to: ", err.Error())
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(data))
	return
}

func (th TaskHandlers) CreateTask(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodPost{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var task domain.Task
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Cannot decode request body due to: ", err.Error())
	}

	task, err = th.storage.CreateTask(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Cannot create task due to: ", err.Error())
	}

	data, err := json.Marshal(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Cannot marshal tasks due to: ", err.Error())
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(data))
	return
}
