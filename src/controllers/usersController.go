package controllers

import (
	"connectopia-api/src/database"
	"connectopia-api/src/models"
	responses "connectopia-api/src/reponses"
	"connectopia-api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetUser retrieves all users.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users!"))
}

// GetUserByID retrieves a user by their ID.
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting user by ID!"))
}

// CreateUser handles the creation of a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
	}

	var userModel models.UserModel
	if err = json.Unmarshal(bodyRequest, &userModel); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
	}

	if err = userModel.Prepare("insert"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	userModel.ID, err = repository.Create((userModel))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}

	responses.JSON(w, http.StatusCreated, userModel)
}

// UpdateUser updates an existing user.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var userModel models.UserModel
	if err = json.Unmarshal(bodyRequest, &userModel); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = userModel.Prepare("update"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if err = repository.Update(ID, userModel); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser deletes a user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user!"))
}
