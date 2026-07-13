package handler

import (
	"AssetFlow/models"
	"AssetFlow/service"
	"AssetFlow/utils"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var user models.RegisterUser

	if err := utils.ParseBody(r, &user); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "failed to parse request body")
		return
	}

	token, statusCode, err := service.RegisterUser(user)
	if err != nil {
		utils.RespondError(w, statusCode, err, "failed to register user")
		return
	}

	utils.RespondJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}{
		Message: "User register successfully",
		Token:   token,
	})

}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.LoginRequest

	if err := utils.ParseBody(r, &user); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "failed to parse request body")
	}

	token, statusCode, err := service.LoginUser(user)
	if err != nil {
		utils.RespondError(w, statusCode, err, "failed to login")
		return
	}
	
	utils.RespondJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}{
		Message: "User login successfully",
		Token:   token,
	})
}
