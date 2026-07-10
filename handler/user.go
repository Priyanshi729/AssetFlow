package handler

import (
	"AssetFlow/database/dbhelper"
	"AssetFlow/models"
	"AssetFlow/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var v = validator.New()

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.RegisterUser

	if parseErr := utils.ParseBody(r, &user); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}

	if err := v.Struct(user); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, fmt.Sprintf("invalid validation failed"))
		return
	}

	exists, existsErr := dbhelper.IsUserExists(user.Email)
	if existsErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, existsErr, "failed to check user existence")
		return
	}
	if exists {
		utils.RespondError(w, http.StatusBadRequest, nil, "user already exists")
		return
	}

	hashedPassword, hasErr := utils.HashPassword(user.Password)
	if hasErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, hasErr, "failed to secure password")
		return
	}

	userID, err := dbhelper.CreateUser(user.Name, user.Email, user.PhoneNumber, user.Role, user.Type, hashedPassword)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to create user")
		return
	}

	token, err := utils.GenerateJWT(userID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to generate token")
		return
	}

	utils.RespondJSON(w, http.StatusOK,
		struct {
			Message string `json:"message"`
			Token   string `json:"token"`
		}{Message: "user created successfully",
			Token: token})
}
