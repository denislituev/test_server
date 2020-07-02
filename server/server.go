package server

import (
	"net/http"
	"test_server/server/handlers"
)


type Server struct{
	addr string
	tasksHandler handlers.TaskHandlers
}

func NewServer(addr string, tasksHandler handlers.TaskHandlers) *Server{
	return &Server{
		addr: addr,
		tasksHandler: tasksHandler,
	}
}

// localhost:8080/tasks    GET
// localhost:8080/task    POST

func (s *Server) ConfigureAndRun(){

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", s.tasksHandler.GetAllTasks)
	mux.HandleFunc("/task", s.tasksHandler.CreateTask)

	http.ListenAndServe(s.addr, mux)

}