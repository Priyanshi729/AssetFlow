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

func GetAssets(w http.ResponseWriter, r *http.Request) {

	assets, statusCode, err := service.GetAssets()
	if err != nil {
		utils.RespondError(w, statusCode, err, "failed to get assets")
		return
	}

	utils.RespondJSON(w, http.StatusOK, assets)
}

func GetAssetByID(w http.ResponseWriter, r *http.Request) {

	assetID := r.PathValue("assetID")

	asset, statusCode, err := service.GetAssetByID(assetID)
	if err != nil {
		utils.RespondError(w, statusCode, err, "failed to get asset")
		return
	}

	utils.RespondJSON(w, http.StatusOK, asset)
}

func UpdateAsset(w http.ResponseWriter, r *http.Request) {
	assetID := r.PathValue("assetID")

	var body models.UpdateAssetRequest

	if err := utils.ParseBody(r, &body); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "failed to parse request body")
		return
	}

	statusCode, err := service.UpdateAsset(assetID, body)
	if err != nil {
		utils.RespondError(w, statusCode, err, "failed to update asset")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{
		"message": "Asset updated successfully",
	})
}
