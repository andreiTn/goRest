package router

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"restApi/taskController"
)

func init() {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", taskController.GetTasks).Methods("GET")
	router.HandleFunc("/task/{id}", taskController.GetTask).Methods("GET")
	router.HandleFunc("/task/{id}", taskController.AddTask).Methods("POST")
	router.HandleFunc("/task/{id}", taskController.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
