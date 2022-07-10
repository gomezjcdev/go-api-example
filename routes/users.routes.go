package routes

import (
	"encoding/json"
	"github.com/gomezjcdev/go-api-gorilla-orm/db"
	"github.com/gomezjcdev/go-api-gorilla-orm/models"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	} else {
		db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
		json.NewEncoder(w).Encode(&user)
	}
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	createdUser := db.DB.Create(&user)

	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else {
		json.NewEncoder(w).Encode(&user)
	}
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	} else {
		db.DB.Delete(&user)
		w.WriteHeader(http.StatusOK)
	}
}
