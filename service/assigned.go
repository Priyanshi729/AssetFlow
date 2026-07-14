package service

import (
	"AssetFlow/database"
	"AssetFlow/models"
	"AssetFlow/repository"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func AssignAsset(assetID string, body models.AssignAssetRequest) (int, error) {
	v := validator.New()

	if err := v.Struct(body); err != nil {
		return http.StatusBadRequest, err
	}

	err := repository.CreateAssetAssignment(database.DB, assetID, body.UserID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = repository.UpdateAssetStatus(assetID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
