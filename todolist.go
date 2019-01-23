package main

import (
	"fmt"
	"net/http"
	"project/dao"
	"project/web"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome !")
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		taskController.GetTask(w, r)
	case "PUT":
		taskController.CreateTask(w, r)
	case "POST":
		taskController.UpdateTask(w, r)
	case "DELETE":
		taskController.DeleteTask(w, r)
	default:
		fmt.Fprintf(w, "Sorry only GET, POST, PUT and DELETE methods are supported.")
	}
}

var taskController *web.TaskController

func main() {
	http.HandleFunc("/", welcomeHandler)

	http.HandleFunc("/tasks", tasksHandler)

	taskDao, err := dao.GetDao(dao.RedisDAO)

	if err != nil {
		fmt.Println(err)
	}

	taskController = web.NewTaskController(taskDao)

	http.ListenAndServe(":8020", nil)
}
