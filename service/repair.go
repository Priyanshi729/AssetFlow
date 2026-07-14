package service

import (
	"AssetFlow/repository"
	"net/http"
)

func SendForRepair(assetID string) (int, error) {
	if err := repository.CreateRepair(assetID); err != nil {
		return http.StatusInternalServerError, err
	}

	if err := repository.UpdateAssetStatus(assetID, "for_repair"); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
