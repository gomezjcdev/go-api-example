package routes

import (
	"encoding/json"
	"github.com/gomezjcdev/go-api-gorilla-orm/db"
	"github.com/gomezjcdev/go-api-gorilla-orm/models"
	"github.com/gorilla/mux"
	"net/http"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var tasks models.Task

	db.DB.First(&tasks, params["id"])

	if tasks.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task Not Found"))
		return
	} else {
		json.NewEncoder(w).Encode(&tasks)
	}
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task = models.Task{}
	json.NewDecoder(r.Body).Decode(&task)
	createdTask := db.DB.Create(&task)

	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else {
		json.NewEncoder(w).Encode(&task)
	}
}

func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	} else {
		db.DB.Delete(&task)
		w.WriteHeader(http.StatusOK)
	}
}
