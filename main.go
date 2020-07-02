package main

import (
	"test_server/server"
	"test_server/server/handlers"
	"test_server/storage/memstorage"
)

func main(){

	ms := memstorage.NewMemStorage()
	h := handlers.NewTaskHandlers(ms)

 	srv := server.NewServer(":8080", h)
	srv.ConfigureAndRun()
}
