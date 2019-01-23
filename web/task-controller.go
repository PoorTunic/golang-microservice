package web

import (
	"encoding/json"
	"net/http"
	"project/dao"
	"project/model"
	"time"
)

// TaskController structure is uset to have a DAO
type TaskController struct {
	taskDao dao.TaskDao
}

// NewTaskController is uset to get the dao controller
func NewTaskController(taskDao dao.TaskDao) *TaskController {
	controller := TaskController{
		taskDao,
	}
	return &controller
}

// CreateTask is uset to create a new task
func (ctrl *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var nTask = model.Task{}

	decodeJson := json.NewDecoder(r.Body)

	err := decodeJson.Decode(&nTask)
	if err != nil {
		panic(err)
	}

	nTask.CreationDate = time.Now()
	nTask.ModificationTime = time.Now()
	nTask.Status = model.StatusTodo

	if resTask, err := ctrl.taskDao.Upsert(&nTask); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resTask)
	}
}

func (ctrl *TaskController) GetTask(w http.ResponseWriter, r *http.Request) {
	key, _ := r.URL.Query()["Id"]

	if resTask, err := ctrl.taskDao.Get(key[0]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resTask)
	}
}

func (ctrl *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var nTask = model.Task{}

	taskJson := json.NewDecoder(r.Body)

	err := taskJson.Decode(&nTask)
	if err != nil {
		panic(err)
	}

	nTask.ModificationTime = time.Now()

	if resTask, err := ctrl.taskDao.Upsert(&nTask); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resTask)
	}
}

func (ctrl *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	paramID := r.URL.Query().Get("Id")

	if err := ctrl.taskDao.Delete(paramID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(true)
	}
}

// -- OLD CONTROLLER

// func CreateTask(w http.ResponseWriter, r *http.Request){
// 	decode := json.NewDecoder(r.Body)

// 	var nTask model.Task

// 	err := decode.Decode(&nTask)
// if err != nil{
// panic(err)
// }

// nTask.CreationDate = time.Now()
// nTask.ModificationTime = time.Now()

// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// w.WriteHeader(http.StatusOK)
// if err := json.NewEncoder(w).Encode(nTask); err != nil {
// http.Error(w, err.Error(), http.StatusInternalServerError)
// }
// }

// func GetTask(w http.ResponseWriter, r *http.Request){
// decode := json.NewDecoder(r.Body)

// var gTask model.Task

// err := decode.Decode(&gTask)
// if err != nil{
// panic(err)
// }

// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// w.WriteHeader(http.StatusOK)
// if err := json.NewEncoder(w).Encode(gTask); err != nil {
// http.Error(w, err.Error(), http.StatusInternalServerError)
// }
// }

// func UpdateTask(w http.ResponseWriter, r *http.Request){
// decode := json.NewDecoder(r.Body)

// var uTask model.Task

// err := decode.Decode(&uTask)
// if err != nil{
// panic(err)
// }

// uTask.ModificationTime = time.Now()

// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// w.WriteHeader(http.StatusOK)
// if err := json.NewEncoder(w).Encode(uTask); err != nil {
// http.Error(w, err.Error(), http.StatusInternalServerError)
// }
// }

// func DeleteTask(w http.ResponseWriter, r *http.Request){
// decode := json.NewDecoder(r.Body)

// var dTask model.Task

// err := decode.Decode(&dTask)
// if err != nil{
// panic(err)
// }

// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// w.WriteHeader(http.StatusOK)
// if err := json.NewEncoder(w).Encode(dTask); err != nil {
// http.Error(w, err.Error(), http.StatusInternalServerError)
// }
// }
