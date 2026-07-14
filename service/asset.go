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

		assetID, err = repository.CreateAsset(tx, asset.Brand, asset.Model, asset.SerialNumber, asset.AssetType, asset.Status, asset.OwnerType, asset.WarrantyStart, asset.WarrantyEnd)
		if err != nil {
			return err
		}

		switch asset.AssetType {

		case "laptop":
			return repository.CreateLaptop(tx, assetID, asset.Processor, asset.RAM, asset.Storage, asset.OperatingSystem, asset.Charger, asset.DevicePassword)

		case "mobile":
			return repository.CreateMobile(tx, assetID, asset.OperatingSystem, asset.RAM, asset.Storage, asset.Charger, asset.DevicePassword)

		case "keyboard":
			return repository.CreateKeyboard(tx, assetID, asset.Layout, asset.Connectivity)

		case "mouse":
			return repository.CreateMouse(tx, assetID, asset.DPI, asset.Connectivity)

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

func GetAssetByID(assetID string) (*models.AssetDetails, int, error) {

	asset, err := repository.GetAssetByID(assetID)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	return asset, http.StatusOK, nil
}
