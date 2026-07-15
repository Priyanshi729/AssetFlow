package service

import (
	"AssetFlow/models"
	"AssetFlow/repository"
	"AssetFlow/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func RegisterUser(user models.RegisterUser) (string, int, error) {

	v := validator.New()
	if err := v.Struct(user); err != nil {
		return "", http.StatusBadRequest, err
	}

	exists, existerr := repository.IsUserExists(user.Email)
	if existerr != nil {
		return "", http.StatusInternalServerError, existerr
	}
	if exists {
		return "", http.StatusBadRequest, nil
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	userID, userErr := repository.CreateUser(user.Name, user.Email, hashedPassword, user.PhoneNumber, user.Role, user.UserType)
	if userErr != nil {
		return "", http.StatusInternalServerError, userErr
	}

	token, tokenerr := utils.GenerateJWT(userID, user.Role)
	if tokenerr != nil {
		return "", http.StatusInternalServerError, tokenerr
	}

	return token, http.StatusOK, nil
}

func LoginUser(user models.LoginRequest) (string, int, error) {

	v := validator.New()

	if err := v.Struct(user); err != nil {
		return "", http.StatusBadRequest, err
	}

	loginData, err := repository.GetUserByPassword(user.Email, user.Password)
	if err != nil {
		return "", http.StatusUnauthorized, err
	}

	token, err := utils.GenerateJWT(loginData.UserID, loginData.Role)

	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return token, http.StatusOK, nil
}

func GetUser(userID string) (*models.User, error) {
	return repository.GetUser(userID)
}

func DeleteUser(userID string) (int, error) {

	if err := repository.DeleteUser(userID); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
