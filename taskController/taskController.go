package taskController

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"restApi/taskModel"
	"restApi/apiMessages"
)

var task taskModel.Task
var apiMess apiMessages.ApiMessage

var tasks []taskModel.Task // get from DB?

func init() {
	tasks = append(tasks, taskModel.Task{ID: "1", Todo: "todo 1", User: "user one"})
}
/**
/ 	@route /tasks
/	- get all tasks
*/
func GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

/**
/ 	@route /task/{id}
/	- get a task based on id
/	- returns task or a "Task not found message"
*/
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var task taskModel.Task
	var found bool

	for _, item := range tasks {
		if item.ID == params["id"] {
			task = item
			found = true
		}
	}

	if found == false {
		apiMess.Message = "Task not found"
		json.NewEncoder(w).Encode(apiMess)
		return
	}
	json.NewEncoder(w).Encode(task)
}

/**
/ 	@route /tasks/{id}
/	- add a task
/	- returns task added message
*/
func AddTask(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)

	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = params["id"]

	task.Add(&tasks)

	apiMess.Message = "Added task " + params["id"]
	json.NewEncoder(w).Encode(apiMess)
}

/**
/ 	@route /tasks
/	- delete task
/	- returns task deleted message
*/
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	task.Remove(id, &tasks)
	apiMess.Message = "Deleted task with id: " + id
	json.NewEncoder(w).Encode(apiMess)
}