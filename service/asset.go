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

func GetAssetByID(assetID string) (interface{}, int, error) {
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
		return struct {
			Asset  *models.Asset
			Laptop *models.LaptopRequestSpecific
		}{
			Asset:  asset,
			Laptop: laptop,
		}, http.StatusOK, nil

	case "mobile":

		mobile, err := repository.GetMobileByID(assetID)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return struct {
			Asset  *models.Asset
			Mobile *models.MobileRequestSpecific
		}{
			Asset:  asset,
			Mobile: mobile,
		}, http.StatusOK, nil

	case "keyboard":

		keyboard, err := repository.GetKeyboardByID(assetID)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return struct {
			Asset    *models.Asset
			Keyboard *models.KeyboardRequestSpecific
		}{
			Asset:    asset,
			Keyboard: keyboard,
		}, http.StatusOK, nil

	case "mouse":

		mouse, err := repository.GetMouseByID(assetID)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return struct {
			Asset *models.Asset
			Mouse *models.MouseRequestSpecific
		}{
			Asset: asset,
			Mouse: mouse,
		}, http.StatusOK, nil

	default:
		return nil, http.StatusBadRequest, fmt.Errorf("invalid asset type")
	}
}

func UpdateAsset(assetID string, body models.UpdateAssetRequest) (int, error) {
	v := validator.New()
	if err := v.Struct(body); err != nil {
		return http.StatusBadRequest, err
	}

	assetType, err := repository.GetAssetType(assetID)
	if err != nil {
		return http.StatusNotFound, err
	}

	err = database.Tx(func(tx *sqlx.Tx) error {
		if err := repository.UpdateAsset(tx, assetID, body.Brand, body.Model, body.SerialNumber, body.Status, body.OwnerType, body.WarrantyStart, body.WarrantyEnd); err != nil {
			return err
		}

		switch assetType {

		case "laptop":

			return repository.UpdateLaptop(tx, assetID, body.Processor, body.RAM, body.Storage, body.OperatingSystem, body.Charger, body.DevicePassword)

		case "mobile":

			return repository.UpdateMobile(tx, assetID, body.OperatingSystem, body.RAM, body.Storage, body.Charger, body.DevicePassword)

		case "keyboard":

			return repository.UpdateKeyboard(tx, assetID, body.Layout, body.Connectivity)

		case "mouse":

			return repository.UpdateMouse(tx, assetID, body.DPI, body.Connectivity)

		default:
			return fmt.Errorf("unsupported asset type")
		}
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
