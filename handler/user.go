package handler

import (
	"AssetFlow/middleware"
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

func GetUser(w http.ResponseWriter, r *http.Request) {

	userCtx := middleware.UserContext(r)
	if userCtx == nil {
		utils.RespondError(w, http.StatusUnauthorized, nil, "unauthorized")
		return
	}

	userID := userCtx.UserID

	user, err := service.GetUser(userID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to get user")
		return
	}

	utils.RespondJSON(w, http.StatusOK, user)
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {

	statusCode, err := service.LogoutUser()
	if err != nil {
		utils.RespondError(w, statusCode, err, "failed to logout")
		return
	}

	utils.RespondJSON(w, statusCode, struct {
		Message string `json:"message"`
	}{
		Message: "User logged out successfully",
	})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	userCtx := middleware.UserContext(r)
	if userCtx == nil {
		utils.RespondError(w, http.StatusUnauthorized, nil, "unauthorized")
		return
	}

	statusCode, err := service.DeleteUser(userCtx.UserID)
	if err != nil {
		utils.RespondError(w, statusCode, err, "failed to delete user")
		return
	}

	utils.RespondJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "user deleted successfully",
	})
}
