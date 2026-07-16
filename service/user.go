package service

import (
	"AssetFlow/models"
	"AssetFlow/repository"
	"AssetFlow/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func RegisterUser(user models.RegisterUser) (string, int, error) {

	v := validator.New()
	if err := v.Struct(user); err != nil {
		return "", http.StatusBadRequest, err
	}

	exists, existErr := repository.IsUserExists(user.Email)
	if existErr != nil {
		return "", http.StatusInternalServerError, existErr
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

	token, tokenErr := utils.GenerateJWT(userID, user.Role)
	if tokenErr != nil {
		return "", http.StatusInternalServerError, tokenErr
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

func GetUserAssets(userID string) ([]models.Asset, int, error) {
	assets, err := repository.GetUserAssets(userID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return assets, http.StatusOK, nil
}

func GetUserAssetByID(userID, assetID string) (*models.Asset, int, error) {

	asset, err := repository.GetUserAssetByID(userID, assetID)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	switch asset.AssetType {
	case "laptop":
		laptop, err := repository.GetLaptopByID(assetID)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
		asset.Laptop = laptop

	case "mobile":
		mobile, err := repository.GetMobileByID(assetID)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
		asset.Mobile = mobile

	case "keyboard":
		keyboard, err := repository.GetKeyboardByID(assetID)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
		asset.Keyboard = keyboard

	case "mouse":
		mouse, err := repository.GetMouseByID(assetID)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
		asset.Mouse = mouse

	default:
		return nil, http.StatusBadRequest, fmt.Errorf("invalid asset type")
	}

	return asset, http.StatusOK, nil
}

func LogoutUser() (int, error) {
	return http.StatusOK, nil
}

func DeleteUser(userID string) (int, error) {

	if err := repository.DeleteUser(userID); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
