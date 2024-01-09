package controllers

import (
	"connectopia-api/src/auth"
	"connectopia-api/src/database"
	"connectopia-api/src/models"
	responses "connectopia-api/src/reponses"
	"connectopia-api/src/repositories"
	"connectopia-api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
	}

	var userModel models.UserModel
	if err = json.Unmarshal(bodyRequest, &userModel); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	databaseUser, err := repository.FindByEmail(userModel.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CompareHashAndPassword(databaseUser.Password, userModel.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(databaseUser.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))
}
