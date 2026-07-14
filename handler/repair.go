package handler

import (
	"AssetFlow/service"
	"AssetFlow/utils"
	"net/http"
)

func SendForRepair(w http.ResponseWriter, r *http.Request) {

	assetID := r.PathValue("assetID")

	statusCode, err := service.SendForRepair(assetID)
	if err != nil {
		utils.RespondError(w, statusCode, err, "failed to send asset for repair")
		return
	}

	utils.RespondJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "Asset send for repair successfully",
	})
}
