package service

import (
	"AssetFlow/database"
	"AssetFlow/models"
	"AssetFlow/repository"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

func CreateAsset(asset models.CreateAssetRequest) (string, int, error) {

	v := validator.New()

	if err := v.Struct(asset); err != nil {
		return "", http.StatusBadRequest, err
	}

	var assetID string

	err := database.Tx(func(tx *sqlx.Tx) error {

		var err error

		assetID, err = repository.CreateAsset(tx, asset)
		if err != nil {
			return err
		}

		switch asset.AssetType {

		case "laptop":
			return repository.CreateLaptop(tx, assetID, asset.Laptop)

		case "mobile":
			return repository.CreateMobile(tx, assetID, asset.Mobile)

		case "keyboard":
			return repository.CreateKeyboard(tx, assetID, asset.Keyboard)

		case "mouse":
			return repository.CreateMouse(tx, assetID, asset.Mouse)

		default:
			return fmt.Errorf("invalid asset type")
		}
	})

	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return assetID, http.StatusCreated, nil
}

func GetAssets() ([]models.Asset, int, error) {

	assets, err := repository.GetAssets()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return assets, http.StatusOK, nil
}

func GetAssetByID(assetID string) (*models.Asset, int, error) {
	asset, err := repository.GetAssetByID(assetID)
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
func UpdateAsset(assetID string, req models.UpdateAssetRequest) (int, error) {
	asset, err := repository.GetAssetByID(assetID)
	if err != nil {
		return http.StatusNotFound, err
	}

	err = database.Tx(func(tx *sqlx.Tx) error {
		if err := repository.UpdateAsset(tx, assetID, req); err != nil {
			return err
		}
		switch asset.AssetType {

		case "laptop":
			if req.Laptop != nil {
				return repository.UpdateLaptop(tx, assetID, req.Laptop)
			}

		case "mobile":
			if req.Mobile != nil {
				return repository.UpdateMobile(tx, assetID, req.Mobile)
			}

		case "keyboard":
			if req.Keyboard != nil {
				return repository.UpdateKeyboard(tx, assetID, req.Keyboard)
			}

		case "mouse":
			if req.Mouse != nil {
				return repository.UpdateMouse(tx, assetID, req.Mouse)
			}

		default:
			return fmt.Errorf("invalid asset type")
		}

		return nil
	})

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func DeleteAsset(assetID string) (int, error) {
	err := repository.DeleteAsset(assetID)
	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}
