package handler

import (
	"AssetFlow/models"
	"AssetFlow/service"
	"AssetFlow/utils"
	"net/http"
)

func CreateAsset(w http.ResponseWriter, r *http.Request) {

	var asset models.CreateAssetRequest

	if err := utils.ParseBody(r, &asset); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "failed to parse request body")
		return
	}

	assetID, statusCode, err := service.CreateAsset(asset)
	if err != nil {
		utils.RespondError(w, statusCode, err, "failed to create asset")
		return
	}

	utils.RespondJSON(w, http.StatusCreated, struct {
		Message string `json:"message"`
		AssetID string `json:"asset_id"`
	}{
		Message: "Asset created successfully",
		AssetID: assetID,
	})
}
