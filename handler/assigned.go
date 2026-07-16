package handler

import (
	"AssetFlow/models"
	"AssetFlow/service"
	"AssetFlow/utils"
	"net/http"
)

func AssignAsset(w http.ResponseWriter, r *http.Request) {
	assetID := r.PathValue("assetID")

	var body models.AssignAssetRequest

	if err := utils.ParseBody(r, &body); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "failed to parse request body")
		return
	}

	statusCode, err := service.AssignAsset(assetID, body)

	if err != nil {
		utils.RespondError(w, statusCode, err, "failed to assign asset")
		return
	}

	utils.RespondJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "Asset assigned successfully",
	})
}

func GetAllAssetAssignmentHistory(w http.ResponseWriter, r *http.Request) {

	history, statusCode, err := service.GetAllAssetAssignmentHistory()
	if err != nil {
		utils.RespondError(w, statusCode, err, "failed to get asset assignment history")
		return
	}

	utils.RespondJSON(w, http.StatusOK, history)
}

func GetAllAssignedAssets(w http.ResponseWriter, r *http.Request) {

	history, statusCode, err := service.GetAllAssignedAsset()
	if err != nil {
		utils.RespondError(w, statusCode, err, "failed to get all asset assignment")
		return
	}

	utils.RespondJSON(w, http.StatusOK, history)
}

func ReturnAsset(w http.ResponseWriter, r *http.Request) {
	assetID := r.PathValue("assetID")

	statusCode, err := service.ReturnAsset(assetID)
	if err != nil {
		utils.RespondError(w, statusCode, err, "failed to return asset")
		return
	}

	utils.RespondJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "Asset returned successfully",
	})
}
